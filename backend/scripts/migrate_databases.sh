#!/bin/bash

export $(grep -v '^#' .env | sed 's/"//g' | xargs)
export PGPASSWORD=$DB_PASSWORD

echo "Running migrations..."

for file in internal/env/migrations/*.sql; do
    echo "Applying: $file"
    psql -U $DB_USER -h $DB_HOST -p $DB_PORT -d $DB_NAME -f $file
    echo "Done: $file"
done

echo "All migrations applied."