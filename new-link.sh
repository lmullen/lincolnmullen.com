#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

date=`date +"%Y-%m-%d-%H-%M-%S"`

# Use Hugo archetypes to do the heavy lifting in creating the new file
hugo new --kind link links/$date.md
# Open the resulting file in VS Code
code "content/links/$date.md"

