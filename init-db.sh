#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    ALTER SYSTEM SET max_connections = '1000';
    ALTER SYSTEM SET shared_buffers = '256MB';
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL 