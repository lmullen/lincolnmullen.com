#!/bin/sh

# Make a new post and open it in VS Code

if [ -z "$1" ]
then
  echo "Provide a slug."
  exit 1
fi

date=`date "+%Y-%m-%d"`
name=$1
# Use Hugo archetypes to do the heavy lifting in creating the new file
hugo new --kind blog blog/$date-$name
# Open the resulting file in VS Code
code "content/blog/$date-$name/index.md"

