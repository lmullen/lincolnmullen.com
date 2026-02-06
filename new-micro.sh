#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

timestamp=$(date +"%Y-%m-%d-%H-%M-%S")
post_file="content/micro/$timestamp.md"

# Prompt for image inclusion
read -p "Include an image? (y/n): " include_image

if [[ "$include_image" =~ ^[Yy]$ ]]; then
    # Get image path (drag-and-drop friendly)
    read -p "Image path (drag and drop file): " image_path
    # Remove quotes and trailing spaces from drag-and-drop
    image_path="${image_path//\'/}"
    image_path="${image_path%% }"

    # Validate image exists
    if [[ ! -f "$image_path" ]]; then
        echo "Error: Image file not found: $image_path"
        exit 1
    fi

    # Get alt text
    read -p "Alt text for image: " alt_text

    # Get file extension (lowercase)
    ext="${image_path##*.}"
    ext=$(echo "$ext" | tr '[:upper:]' '[:lower:]')

    # Copy image to assets directory with timestamp name
    new_image_name="$timestamp.$ext"
    cp "$image_path" "assets/images/micro/$new_image_name"
    echo "Image copied to assets/images/micro/$new_image_name"
fi

# Create the micro post using Hugo archetype
hugo new --kind micro "micro/$timestamp.md"

# If image was included, append shortcode to post
if [[ "$include_image" =~ ^[Yy]$ ]]; then
    echo "" >> "$post_file"
    echo "{{< image src=\"/images/micro/$new_image_name\" alt=\"$alt_text\" >}}" >> "$post_file"
fi

# Open in VS Code
code "$post_file"
