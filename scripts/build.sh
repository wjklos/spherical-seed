#!/usr/bin/env bash

dep ensure

echo "Compiling functions to bin/handlers/ ..."

rm -rf bin/

cd src/handlers/
for f in */; do
  dirname="${f%?}"
  filename="${dirname}.go"
  if GOOS=linux go build -o "../../bin/handlers/$dirname" $dirname/${filename}; then
    echo "✓ Compiled $filename"
  else
    echo "✕ Failed to compile $filename!"
    exit 1
  fi
done

echo "Done."