#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

# Make a new post and open it in VS Code

if [ -z "$1" ]
then
  echo "Provide a title."
  exit 1
fi

# Title will be provided by VS Code as an argument
title=$1
# The slug will be generated from the title
slug=$(echo "$title" | tr '[:upper:]' '[:lower:]' | tr -d '[:punct:]' | tr '[:blank:]' '-' | tr -s '-')
date=`date "+%Y-%m-%d"`
# Use Hugo archetypes to do the heavy lifting in creating the new file
POST_TITLE="$title" POST_SLUG="$slug" hugo new --kind blog blog/$date-$slug
# Open the resulting file in VS Code
code "content/blog/$date-$slug/index.md"

