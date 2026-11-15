#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

# The build command for Cloudflare Workers Builds (see wrangler.jsonc).
#
# Cloudflare's build image has an older Hugo and no mise, so this script fetches
# the Hugo that mise would install: the version pinned in .mise.toml, from the
# URL recorded in mise.lock, verified against the checksum recorded there. To
# upgrade Hugo, change .mise.toml and run `mise install` (or `mise lock`); this
# script needs no edit.
#
# Usage: ./cloudflare-build.sh [extra hugo flags]

cd "$(dirname "$0")"

case "$(uname -s)-$(uname -m)" in
  Linux-x86_64) platform=linux-x64 ;;
  Linux-aarch64 | Linux-arm64) platform=linux-arm64 ;;
  *) echo "cloudflare-build: unsupported platform $(uname -s)-$(uname -m)" >&2; exit 1 ;;
esac

pinned=$(sed -n 's/^hugo-extended *= *"\([^"]*\)".*/\1/p' .mise.toml)

# Print a field from the mise.lock table whose header line is exactly $1.
lockfield() { # header field
  awk -v header="$1" -v field="$2" '
    /^\[/ { in_table = ($0 == header) }
    in_table && $1 == field { gsub(/"/, "", $3); print $3; exit }
  ' mise.lock
}

locked=$(lockfield '[[tools.hugo-extended]]' version)
url=$(lockfield "[tools.hugo-extended.\"platforms.$platform\"]" url)
checksum=$(lockfield "[tools.hugo-extended.\"platforms.$platform\"]" checksum)
checksum=${checksum#sha256:}

if [ -z "$pinned" ] || [ "$pinned" != "$locked" ]; then
  echo "cloudflare-build: .mise.toml pins Hugo ${pinned:-(none)} but mise.lock has ${locked:-(none)}; run 'mise lock' and commit mise.lock" >&2
  exit 1
fi
if [ -z "$url" ] || [ -z "$checksum" ]; then
  echo "cloudflare-build: mise.lock has no Hugo download for $platform; run 'mise lock' and commit mise.lock" >&2
  exit 1
fi

work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT
curl --fail --silent --show-error --location --retry 3 --output "$work/hugo.tar.gz" "$url"
echo "$checksum  $work/hugo.tar.gz" | sha256sum --check --quiet
tar -xzf "$work/hugo.tar.gz" -C "$work" hugo
"$work/hugo" version

"$work/hugo" --cleanDestinationDir --minify --baseURL "https://lincolnmullen.com" "$@"
