#!/usr/bin/env bash
set -euo pipefail

BASE_URL="https://lincolnmullen.com"
BUTTONDOWN_API="https://api.buttondown.com/v1"
DRY_RUN=false

usage() {
    echo "Usage: $0 [--dry-run] <path-to-newsletter-index.md>"
    echo ""
    echo "Upload a newsletter to Buttondown as a draft email."
    echo ""
    echo "Options:"
    echo "  --dry-run  Print transformed markdown and API calls without uploading"
    exit 1
}

# Parse arguments
while [[ $# -gt 0 ]]; do
    case "$1" in
        --dry-run)
            DRY_RUN=true
            shift
            ;;
        -h|--help)
            usage
            ;;
        *)
            INPUT_FILE="$1"
            shift
            ;;
    esac
done

if [[ -z "${INPUT_FILE:-}" ]]; then
    usage
fi

if [[ ! -f "$INPUT_FILE" ]]; then
    echo "Error: File not found: $INPUT_FILE" >&2
    exit 1
fi

if [[ "$DRY_RUN" == false ]] && [[ -z "${BUTTONDOWN_HUGO:-}" ]]; then
    echo "Error: BUTTONDOWN_HUGO environment variable is not set" >&2
    exit 1
fi

NEWSLETTER_DIR="$(dirname "$INPUT_FILE")"

# Extract front matter fields
TITLE="$(awk '/^---$/{n++; next} n==1 && /^title:/{sub(/^title: *"?/, ""); sub(/"? *$/, ""); print}' "$INPUT_FILE")"
SLUG="$(awk '/^---$/{n++; next} n==1 && /^slug:/{sub(/^slug: *"?/, ""); sub(/"? *$/, ""); print}' "$INPUT_FILE")"

if [[ -z "$TITLE" ]]; then
    echo "Error: Could not extract title from front matter" >&2
    exit 1
fi
if [[ -z "$SLUG" ]]; then
    echo "Error: Could not extract slug from front matter" >&2
    exit 1
fi

SUBJECT="Working on It ${SLUG}: ${TITLE}"

echo "Subject: $SUBJECT"
echo "Directory: $NEWSLETTER_DIR"
echo ""

# Extract body (everything after second ---)
BODY="$(awk '/^---$/{n++; next} n>=2{print}' "$INPUT_FILE")"

# Strip <!--more--> comments
BODY="$(printf '%s\n' "$BODY" | sed 's/<!--more-->//')"

# Convert relative markdown links to absolute
# Matches ](/ but not ](http or ](// (protocol-relative)
BODY="$(printf '%s\n' "$BODY" | sed -E "s|\]\(/([^)]*)\)|\](${BASE_URL}/\1)|g")"

# Convert em dashes and en dashes (--- → —, -- → –)
# Process longer sequences first to avoid double-conversion
BODY="$(printf '%s\n' "$BODY" | sed 's/---/—/g' | sed 's/--/–/g')"

# Add extra blank lines before ### headings for email spacing
BODY="$(printf '%s\n' "$BODY" | sed -E 's/^### /\
\
\
### /')"

# Upload an image to Buttondown and print the hosted URL
upload_image() {
    local filepath="$1"

    if [[ "$DRY_RUN" == true ]]; then
        echo "[DRY-RUN uploaded: $(basename "$filepath")]"
        return
    fi

    local response
    response="$(curl -s -X POST "${BUTTONDOWN_API}/images" \
        -H "Authorization: Token ${BUTTONDOWN_HUGO}" \
        -F "image=@${filepath}")"

    # Try to extract the image URL from the response
    local url
    url="$(echo "$response" | jq -r '.image // .url // .id // empty' 2>/dev/null || true)"

    if [[ -z "$url" ]]; then
        echo "Error: Failed to upload image $(basename "$filepath")" >&2
        echo "Response: $response" >&2
        exit 1
    fi

    echo "$url"
}

# Process image shortcodes
# We process line by line to handle shortcode replacement robustly
process_shortcodes() {
    local body="$1"

    while IFS= read -r line || [[ -n "$line" ]]; do
        # Check if line contains an image or figure shortcode
        if printf '%s\n' "$line" | grep -qE '\{\{<\s*(image|figure)\s'; then
            local shortcode_type
            shortcode_type="$(printf '%s\n' "$line" | sed -E 's/.*\{\{<[[:space:]]*(image|figure)[[:space:]].*/\1/')"

            # Extract src parameter
            local src=""
            src="$(printf '%s\n' "$line" | sed -E 's/.*src="([^"]+)".*/\1/')"

            # Extract alt parameter (for image shortcode)
            local alt=""
            if printf '%s\n' "$line" | grep -q 'alt="'; then
                alt="$(printf '%s\n' "$line" | sed -E 's/.*alt="([^"]+)".*/\1/')"
            fi

            # Extract caption parameter (for figure shortcode)
            local caption=""
            if printf '%s\n' "$line" | grep -q 'caption="'; then
                caption="$(printf '%s\n' "$line" | sed -E 's/.*caption="([^"]+)".*/\1/')"
            fi

            # Extract link parameter
            local link=""
            if printf '%s\n' "$line" | grep -q 'link="'; then
                link="$(printf '%s\n' "$line" | sed -E 's/.*link="([^"]+)".*/\1/')"
            fi

            # Use caption as alt text for figure shortcodes
            if [[ "$shortcode_type" == "figure" ]] && [[ -n "$caption" ]]; then
                alt="$caption"
            fi

            # Resolve image file path
            local img_path=""
            if [[ "$src" == /* ]]; then
                # Absolute path — look in assets directory
                img_path="${NEWSLETTER_DIR}/../../assets${src}"
            else
                # Relative path — page resource in newsletter directory
                img_path="${NEWSLETTER_DIR}/${src}"
            fi

            if [[ ! -f "$img_path" ]]; then
                echo "Warning: Image not found: $img_path" >&2
                printf '%s\n' "$line"
                continue
            fi

            # Convert SVG to PNG if needed
            local upload_path="$img_path"
            local ext="${src##*.}"
            if [[ "$(echo "$ext" | tr '[:upper:]' '[:lower:]')" == "svg" ]]; then
                local png_path="${img_path%.svg}.png"
                if [[ "$DRY_RUN" == true ]]; then
                    echo "[DRY-RUN] Would convert $(basename "$img_path") to PNG" >&2
                else
                    if ! command -v rsvg-convert &>/dev/null; then
                        echo "Error: rsvg-convert not found. Install with: brew install librsvg" >&2
                        exit 1
                    fi
                    echo "Converting SVG to PNG: $(basename "$img_path")" >&2
                    rsvg-convert -o "$png_path" "$img_path"
                fi
                upload_path="$png_path"
            fi

            # Upload image
            echo "Uploading: $(basename "$upload_path")" >&2
            local hosted_url
            hosted_url="$(upload_image "$upload_path")"

            # Clean up converted PNG
            if [[ "$(echo "$ext" | tr '[:upper:]' '[:lower:]')" == "svg" ]] && [[ "$DRY_RUN" == false ]] && [[ -f "${img_path%.svg}.png" ]]; then
                rm "${img_path%.svg}.png"
            fi

            # Build markdown replacement
            local md_img="![${alt}](${hosted_url})"
            if [[ -n "$link" ]]; then
                md_img="[![${alt}](${hosted_url})](${link})"
            fi

            printf '%s\n' "$md_img"
        else
            printf '%s\n' "$line"
        fi
    done <<< "$body"
}

echo "Processing shortcodes..."
BODY="$(process_shortcodes "$BODY")"

if [[ "$DRY_RUN" == true ]]; then
    echo ""
    echo "=== Transformed Markdown ==="
    echo "$BODY"
    echo ""
    echo "=== API Call ==="
    echo "POST ${BUTTONDOWN_API}/emails"
    echo "Subject: $SUBJECT"
    echo "Status: draft"
    exit 0
fi

# Create draft email on Buttondown
echo ""
echo "Creating draft on Buttondown..."

BODY_TMPFILE="$(mktemp)"
printf '%s\n' "$BODY" > "$BODY_TMPFILE"

RESPONSE="$(jq -n \
    --arg subject "$SUBJECT" \
    --rawfile body "$BODY_TMPFILE" \
    '{subject: $subject, body: $body, status: "draft"}' \
    | curl -s -X POST "${BUTTONDOWN_API}/emails" \
        -H "Authorization: Token ${BUTTONDOWN_HUGO}" \
        -H "Content-Type: application/json" \
        -d @-)"

rm -f "$BODY_TMPFILE"

# Check for errors
if echo "$RESPONSE" | jq -e '.id' &>/dev/null; then
    DRAFT_ID="$(echo "$RESPONSE" | jq -r '.id')"
    echo "Draft created successfully. ID: $DRAFT_ID"
else
    echo "Error creating draft:" >&2
    echo "$RESPONSE" | jq . 2>/dev/null || echo "$RESPONSE" >&2
    exit 1
fi
