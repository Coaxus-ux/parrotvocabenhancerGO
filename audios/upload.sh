#!/bin/bash

BUCKET_NAME="voices"
FOLDER_PATH="us"

for file in $FOLDER_PATH/*; do
  if [ -f "$file" ]; then
    echo "Uploading $file..."
    npx wrangler r2 object put "$BUCKET_NAME/us/$(basename "$file")" --file "$file"
  fi
done
