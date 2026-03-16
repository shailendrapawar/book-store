#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: ./scripts/create_migration.sh <name>"
  echo "Example: ./scripts/create_migration.sh create_users"
  exit 1
fi

MIGRATIONS_DIR="internal/db/migrations"
mkdir -p $MIGRATIONS_DIR

NEXT=$(ls $MIGRATIONS_DIR/*.sql 2>/dev/null | wc -l)
NEXT=$(( NEXT + 1 ))
SEQUENCE=$(printf "%04d" $NEXT)

touch "$MIGRATIONS_DIR/${SEQUENCE}_create_$1.sql"

echo "Created: ${SEQUENCE}_$1.sql"