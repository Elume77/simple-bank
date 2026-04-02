#!/bin/sh

# Exit immediately if a command exits with a non-zero status
set -e

# Load environment variables from app.env


#echo "Using: $DB_URL"

if [ -z "$DB_URL" ]; then
  echo "Error: DB_URL environment variable is not set."
  exit 1
fi

echo "Running database migrations..."
# Use the environment variables from your app.env/environment
# source /app/app.env    
/usr/local/bin/migrate -path ./db/migration -database "$DB_URL" -verbose up


echo "Starting the application..."
# Exec replaces the shell process with the app process so it receives OS signals (SIGTERM)
#exec ./main

exec "$@"