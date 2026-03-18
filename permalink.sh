#!/bin/bash
# Given a content file path, print and copy its permalink to the clipboard.
set -euo pipefail

file="$1"
rel="${file#*content/}"
type="${rel%%/*}"

if [ "$type" = "micro" ]; then
  slug=$(basename "$file" .md)
elif [ "$type" = "blog" ] || [ "$type" = "newsletter" ]; then
  slug=$(grep -m1 '^slug:' "$file" | sed 's/^slug: *//')
else
  echo "Not a recognized post type" >&2
  exit 1
fi

link="/$type/$slug/"
echo "$link"
printf '%s' "$link" | pbcopy
echo "(copied to clipboard)"
