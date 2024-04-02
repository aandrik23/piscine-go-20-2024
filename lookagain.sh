#!/bin/bash

# Look for all files ending with .sh recursively in the current directory and its sub-folders
sh_files=$(find . -type f -name "*.sh" -exec basename {} \; | sed 's/\.sh$//' | sort -r)

# Display the list of files without the .sh extension in descending order
echo "$sh_files"
