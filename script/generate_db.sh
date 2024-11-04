#!/bin/bash

# Set DSN and schema parameters
DSN="postgresql://postgres:postgres@localhost:5432/security_proof?sslmode=disable"
DB_GENERATE_PATH="/Users/wanho/GolandProjects/security-proof/internal/db"

# Run jet command for user schema
jet -dsn="$DSN" -schema=user -path=$DB_GENERATE_PATH

# Run jet command for proof schema
jet -dsn="$DSN" -schema=proof -path=$DB_GENERATE_PATH

echo "Database schema generation completed for user and proof schemas."
