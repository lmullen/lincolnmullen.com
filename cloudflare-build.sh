#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

# The build command for Cloudflare Workers Builds (see wrangler.jsonc).
#
# Cloudflare's build image has an older Hugo, no Dart Sass, and no mise, so this
# script fetches the tools mise would install: Hugo and Dart Sass at the
# versions pinned in .mise.toml, from the URLs recorded in mise.lock, verified
# against the checksums recorded there. Hugo compiles the SCSS by running Dart
# Sass, which it finds as `sass` on the PATH. To upgrade either tool, change
# .mise.toml and run `mise install` (or `mise lock`); this script needs no edit.
#
# Usage: ./cloudflare-build.sh [extra hugo flags]

cd "$(dirname "$0")"

case "$(uname -s)-$(uname -m)" in
  Linux-x86_64) platform=linux-x64 ;;
  Linux-aarch64 | Linux-arm64) platform=linux-arm64 ;;
  *) echo "cloudflare-build: unsupported platform $(uname -s)-$(uname -m)" >&2; exit 1 ;;
esac

# Print a field from the mise.lock table whose header line is exactly $1.
lockfield() { # header field
  awk -v header="$1" -v field="$2" '
    /^\[/ { in_table = ($0 == header) }
    in_table && $1 == field { gsub(/"/, "", $3); print $3; exit }
  ' mise.lock
}

work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT

# Download the tool whose key in .mise.toml and mise.lock is $2 into
# $work/$1, after checking that mise.lock records the pinned version and that the
# archive matches the recorded checksum.
fetch() { # name lock-key
  local name=$1 key=$2 pinned locked url checksum
  pinned=$(sed -n "s|^${key} *= *\"\([^\"]*\)\".*|\1|p" .mise.toml)
  locked=$(lockfield "[[tools.${key}]]" version)
  url=$(lockfield "[tools.${key}.\"platforms.$platform\"]" url)
  checksum=$(lockfield "[tools.${key}.\"platforms.$platform\"]" checksum)
  checksum=${checksum#sha256:}
  if [ -z "$pinned" ] || [ "$pinned" != "$locked" ]; then
    echo "cloudflare-build: .mise.toml pins $name ${pinned:-(none)} but mise.lock has ${locked:-(none)}; run 'mise lock' and commit mise.lock" >&2
    exit 1
  fi
  if [ -z "$url" ] || [ -z "$checksum" ]; then
    echo "cloudflare-build: mise.lock has no $name download for $platform; run 'mise lock' and commit mise.lock" >&2
    exit 1
  fi
  curl --fail --silent --show-error --location --retry 3 --output "$work/$name.tar.gz" "$url"
  echo "$checksum  $work/$name.tar.gz" | sha256sum --check --quiet
  mkdir "$work/$name"
  tar -xzf "$work/$name.tar.gz" -C "$work/$name"
}

fetch hugo hugo
fetch dart-sass '"github:sass/dart-sass"'

# The Dart Sass archive holds a dart-sass/ directory with the `sass` launcher
# and the runtime it needs; Hugo looks for `sass` on the PATH.
export PATH="$work/dart-sass/dart-sass:$work/hugo:$PATH"
hugo version
sass --version

hugo --cleanDestinationDir --minify --baseURL "https://lincolnmullen.com" "$@"
