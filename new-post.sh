#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

# Create a new blog post. Handles both titled (long-form) and untitled
# (short-form/micro) posts. Title is passed as $1 (blank = untitled).
#
# Usage: ./new-post.sh "Post title"
#        ./new-post.sh ""              # untitled/short-form
#
# Archetype: archetypes/blog/index.md

cd "$(dirname "$0")"

# Capture timestamps up front so filename and front matter always match.
timestamp=$(date +"%Y-%m-%d-%H-%M-%S")
iso_date=$(date +"%Y-%m-%dT%H:%M:%S%z" | sed 's/\(..\)$/:\1/')
date_prefix=$(date +"%Y-%m-%d")

# Title from first argument (blank = untitled short-form post)
title="${1:-}"

if [ -n "$title" ]; then
  slug=$(echo "$title" | tr '[:upper:]' '[:lower:]' | tr -d '[:punct:]' | tr '[:blank:]' '-' | tr -s '-')
  dir_name="$date_prefix-$slug"
else
  slug="$timestamp"
  dir_name="$timestamp"
fi

post_dir="content/blog/$dir_name"
post_file="$post_dir/index.md"

# Create the post using the Hugo archetype
HUGO_POST_TITLE="$title" HUGO_POST_DATE="$iso_date" HUGO_POST_SLUG="$slug" hugo new --kind blog "blog/$dir_name"

# Open in VS Code
code "$post_file"
