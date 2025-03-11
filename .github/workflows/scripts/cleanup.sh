#!/bin/bash

# Gets nth parent directory based on the pwd tree
get_nth_parent_pwd() {
    local n="$1"
    local dir="$PWD"

    for ((i=0; i<n; i++)); do
        dir=$(dirname "$dir")
    done

    echo "$dir"
}

# Define cleanup markers
START_MARKER="> Start clean up"
END_MARKER="> End clean up"

# Find all files (excluding .git and .github directories)
# Assumes that the script is in .github/workflows/scripts directory
parent=$(dirname $PWD)
grandparent=$(dirname $parent)
FILES=$(find $(get_nth_parent_pwd 3) -type f -path "./.git/*" -prune -o -name "*.go")

echo $FILES
# Process each file
for FILE in $FILES; do

    TEMP_FILE="$(mktemp)"
    DELETING_MODE=false
    MODIFIED=false

    # Read file line by line
    while IFS= read -r line; do

        if [[ "$line" == *"$START_MARKER"* ]]; then
            DELETING_MODE=true
            MODIFIED=true
            continue
        fi

        if [[ "$line" == *"$END_MARKER"* ]]; then
            echo "FOUND END MARKER"
            DELETING_MODE=false
            continue
        fi

        if ! $DELETING_MODE; then
            echo "$line" >> "$TEMP_FILE"
        fi

    done < "$FILE"
    
    # Only replace the file if it was modified
    if $MODIFIED; then
        mv "$TEMP_FILE" "$FILE"
        echo "✅ Cleaned: $FILE"
    else
        rm "$TEMP_FILE"
    fi
done

echo "🚀 Cleanup complete for all files!"

