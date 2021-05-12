#!/bin/sh

set -e

echo "run db migrations"
/account/migrate -path /account/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

echo "start app"
exec "$@"
