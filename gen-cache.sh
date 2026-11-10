#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

# Keep Hugo's processed-image cache (resources/_gen/images) exactly in step
# with what a production build needs.
#
# The cache is committed so that Cloudflare Workers Builds, which starts from a
# fresh clone, reuses the encoded images instead of re-encoding all of them on
# every deploy. That only works if the cache holds every image a production
# build processes (anything missing is re-encoded on every deploy) and nothing
# else (stale entries bloat the repository for no benefit).
#
# Usage: ./gen-cache.sh          # sync: add missing variants, remove stale ones
#        ./gen-cache.sh --check  # report drift and exit 1 without changing anything
#
# Both modes build the site with the production settings (no drafts, no future
# or expired posts) into a throwaway directory, using a copy of the cache, and
# run Hugo's garbage collection on that copy. The copy is then the exact set of
# files production needs. Sync mode makes resources/_gen/images match it; check
# mode only compares. Hugo must be the version pinned in .mise.toml, because the
# Cloudflare build uses that version and cache keys can change between versions.

cd "$(dirname "$0")"

mode=sync
case "${1:-}" in
  "") ;;
  --check) mode=check ;;
  *) echo "usage: $0 [--check]" >&2; exit 2 ;;
esac

cache=resources/_gen/images

# The pinned Hugo version is the source of truth for both local dev and the
# Cloudflare build; make sure the two pins agree and that we are running it.
pinned=$(sed -n 's/^hugo-extended *= *"\([^"]*\)".*/\1/p' .mise.toml)
if [ -z "$pinned" ]; then
  echo "gen-cache: could not read the hugo-extended version from .mise.toml" >&2
  exit 1
fi
if ! grep -q "hugo_extended_${pinned}_linux-amd64" wrangler.jsonc; then
  echo "gen-cache: wrangler.jsonc does not download Hugo ${pinned}; keep its build command in sync with .mise.toml" >&2
  exit 1
fi
running=$(hugo version 2>/dev/null | sed -n 's/^hugo v\([0-9][0-9.]*\).*/\1/p')
if [ "$running" != "$pinned" ]; then
  echo "gen-cache: hugo ${running:-is not installed}; .mise.toml pins ${pinned}. Run 'mise install' (or fetch that release from github.com/gohugoio/hugo/releases) so the cache is built with the version Cloudflare uses." >&2
  exit 1
fi
if [ ! -d node_modules ]; then
  echo "gen-cache: node_modules is missing; run 'npm ci' first (the SCSS imports Bootstrap)" >&2
  exit 1
fi

work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT
mkdir -p "$work/resources/_gen" "$work/public"
cp -R "$cache" "$work/resources/_gen/images"

# Build with the production settings, then let --gc drop every cache file the
# build did not use. HUGO_RESOURCEDIR points Hugo at the copy.
HUGO_RESOURCEDIR="$work/resources" hugo --gc --quiet --destination "$work/public"

list() { (cd "$1" && find . -type f | sort); }
missing=$(comm -13 <(list "$cache") <(list "$work/resources/_gen/images"))
stale=$(comm -23 <(list "$cache") <(list "$work/resources/_gen/images"))
n_missing=$(printf '%s' "$missing" | grep -c . || true)
n_stale=$(printf '%s' "$stale" | grep -c . || true)

if [ "$n_missing" -eq 0 ] && [ "$n_stale" -eq 0 ]; then
  echo "gen-cache: $cache is up to date ($(list "$cache" | wc -l | tr -d ' ') files)"
  exit 0
fi

[ "$n_missing" -gt 0 ] && { echo "Missing from the cache ($n_missing); production re-encodes these on every deploy:"; printf '%s\n' "$missing" | sed 's|^\./|  + |'; }
[ "$n_stale" -gt 0 ] && { echo "Stale in the cache ($n_stale); no production page uses these:"; printf '%s\n' "$stale" | sed 's|^\./|  - |'; }

if [ "$mode" = check ]; then
  echo "gen-cache: $cache is out of date; run ./gen-cache.sh and commit the result" >&2
  exit 1
fi

if [ "$n_missing" -gt 0 ]; then
  printf '%s\n' "$missing" | while IFS= read -r f; do
    mkdir -p "$cache/$(dirname "$f")"
    cp "$work/resources/_gen/images/$f" "$cache/$f"
  done
fi
if [ "$n_stale" -gt 0 ]; then
  printf '%s\n' "$stale" | while IFS= read -r f; do
    rm "$cache/$f"
  done
fi
find "$cache" -mindepth 1 -type d -empty -delete
echo "gen-cache: updated $cache (+$n_missing, -$n_stale); commit it with: git add -A $cache"
