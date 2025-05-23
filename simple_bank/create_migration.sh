#!/bin/bash

# Usage: ./create_migration.sh init_schema

# Check if name is passed
if [ -z "$1" ]; then
  echo "Usage: $0 <migration_name>"
  exit 1
fi

MIGRATION_NAME=$1
DIR="db/migration"
EXT="sql"

# Create the migration directory if it doesn't exist
mkdir -p "$DIR"

# Find the highest existing sequence number
LAST_SEQ=$(ls "$DIR"/*.${EXT} 2>/dev/null | \
  grep -oE '[0-9]+' | \
  sort -n | \
  tail -n 1)

# If no files found, start at 1
if [ -z "$LAST_SEQ" ]; then
  NEXT_SEQ=1
else
  NEXT_SEQ=$((LAST_SEQ + 1))
fi

# Pad sequence number to 4 digits
SEQ_PADDED=$(printf "%04d" "$NEXT_SEQ")

# Build filenames
UP_FILE="${DIR}/${SEQ_PADDED}_${MIGRATION_NAME}.up.${EXT}"
DOWN_FILE="${DIR}/${SEQ_PADDED}_${MIGRATION_NAME}.down.${EXT}"

# Create the migration files
echo "-- +migrate Up" > "$UP_FILE"
echo "-- +migrate Down" > "$DOWN_FILE"

# Report result
echo "Created:"
echo " - $UP_FILE"
echo " - $DOWN_FILE"

