#!/bin/sh

# Exit immediately if a command exits with a non-zero status
set -e

echo "Running database migrations..."
# Use the environment variables from your app.env/environment
# Assuming DB_SOURCE is your connection string (e.g., postgresql://user:pass@host:5432/db)
/usr/local/bin/migrate -path ./db/migration -database "$DB_URL" -verbose up

echo "Starting the application..."
# Exec replaces the shell process with the app process so it receives OS signals (SIGTERM)
exec ./main