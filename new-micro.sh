#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

date=$(date +"%Y-%m-%d-%H-%M-%S")

# Use Hugo archetypes to create the new file
hugo new --kind micro micro/$date.md
# Open the resulting file in VS Code
code "content/micro/$date.md"
