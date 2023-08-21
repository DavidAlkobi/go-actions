#!/bin/bash

# Define the directory path where old files are located
TARGET_DIRECTORY="."

# Define the maximum age of files to be considered as "old" (in days)
MAX_AGE_DAYS=30

# Navigate to the target directory
cd "$TARGET_DIRECTORY" || exit 1

# Find and delete files older than MAX_AGE_DAYS
find . -type f -mtime +$MAX_AGE_DAYS -exec rm {} \;

echo "Cleanup operation completed."
