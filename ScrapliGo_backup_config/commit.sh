#!/bin/bash
message="$1"
if [ -z "$message" ]; then
  echo "Usage: $0 <commit message>"
  exit 1
fi
./scrapligobackupconfig
git add .
git commit -m "$message"
git push
