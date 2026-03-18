#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

# Create a new blog post. Handles both titled (long-form) and untitled
# (short-form/micro) posts, with optional image inclusion.

cd "$(dirname "$0")"

# Capture timestamps up front so filename and front matter always match.
timestamp=$(date +"%Y-%m-%d-%H-%M-%S")
iso_date=$(date +"%Y-%m-%dT%H:%M:%S%z" | sed 's/\(..\)$/:\1/')
date_prefix=$(date +"%Y-%m-%d")

# Prompt for title (blank = untitled short-form post)
read -p "Post title (leave blank for untitled): " title

if [ -n "$title" ]; then
  slug=$(echo "$title" | tr '[:upper:]' '[:lower:]' | tr -d '[:punct:]' | tr '[:blank:]' '-' | tr -s '-')
  dir_name="$date_prefix-$slug"
else
  slug="$timestamp"
  dir_name="$timestamp"
fi

post_dir="content/blog/$dir_name"
post_file="$post_dir/index.md"
mkdir -p "$post_dir"

# Prompt for image inclusion
read -p "Include an image? (y/n): " include_image

image_name=""
alt_text=""
if [[ "$include_image" =~ ^[Yy]$ ]]; then
  read -p "Image path (drag and drop file): " image_path
  image_path="${image_path//\'/}"
  image_path="${image_path%% }"

  if [[ ! -f "$image_path" ]]; then
    echo "Error: Image file not found: $image_path"
    exit 1
  fi

  read -p "Alt text for image: " alt_text

  ext="${image_path##*.}"
  ext=$(echo "$ext" | tr '[:upper:]' '[:lower:]')
  image_name="$timestamp.$ext"
  cp "$image_path" "$post_dir/$image_name"
  echo "Image copied to $post_dir/$image_name"
fi

# Write front matter and content
{
  echo "---"
  if [ -n "$title" ]; then
    echo "title: \"$title\""
  fi
  echo "date: \"$iso_date\""
  echo "slug: \"$slug\""
  echo "draft: true"
  if [ -n "$image_name" ]; then
    echo "image: \"$image_name\""
  fi
  echo "description: \"\""
  if [ -n "$title" ]; then
    echo "# tags:"
    echo "# -"
    echo "# crosspost:"
    echo "#   url: \"\""
    echo "#   source: \"\""
  fi
  echo "# linkpost: \"\""
  echo "# via: \"\""
  echo "social:"
  echo "  bluesky: \"\""
  echo "  mastodon: \"\""
  echo "  linkedin: \"\""
  echo "---"
  echo ""
  if [ -n "$image_name" ]; then
    echo "{{< image src=\"$image_name\" alt=\"$alt_text\" >}}"
    echo ""
  fi
} > "$post_file"

# Open in VS Code
code "$post_file"
