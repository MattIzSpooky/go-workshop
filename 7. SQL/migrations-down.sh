#!/bin/bash
# yeah, username = postgres
# password = postgres
POSTGRESQL_URL='postgres://postgres:postgres@localhost:5432/workshop_db?sslmode=disable'

echo "Running migrations DOWN.."

migrate -database "$POSTGRESQL_URL" -path sql/migrations down

echo "done!"