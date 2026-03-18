#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

# Create a new blog post. Handles both titled (long-form) and untitled
# (short-form/micro) posts, with optional image inclusion.
#
# Archetype: archetypes/blog/index.md

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
fi

# Create the post using the Hugo archetype
HUGO_POST_TITLE="$title" HUGO_POST_SLUG="$slug" hugo new --kind blog "blog/$dir_name"

# Overwrite the date so it matches the captured timestamp exactly
sed -i '' "s|^date: .*|date: '$iso_date'|" "$post_file"

# Remove the title line for untitled posts
if [ -z "$title" ]; then
  sed -i '' '/^title: ""$/d' "$post_file"
fi

# Handle image: copy into page bundle and update front matter/content
if [ -n "$image_name" ]; then
  cp "$image_path" "$post_dir/$image_name"
  echo "Image copied to $post_dir/$image_name"

  # Insert image field before description in front matter
  sed -i '' "s|^description:|image: \"$image_name\"\ndescription:|" "$post_file"

  # Append image shortcode to content
  echo "{{< image src=\"$image_name\" alt=\"$alt_text\" >}}" >> "$post_file"
fi

# Open in VS Code
code "$post_file"
