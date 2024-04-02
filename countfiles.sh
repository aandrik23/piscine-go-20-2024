countfiles.sh
#!/bin/bash

# Count the number of regular files and directories in the current directory and its sub-folders
count=$(find . -type f -o -type d | wc -l)

# Print the count
echo "$count"
