#!/usr/bin/env zsh
set -euo pipefail

DB_USER=${DB_USER:-root}
DB_PASS=${DB_PASS:-admin123}
DB_HOST=${DB_HOST:-127.0.0.1}
DB_PORT=${DB_PORT:-3306}
SQL_FILE="$(cd "$(dirname "$0")/.." && pwd)/migrations/mysql/alter_not_null.sql"

if ! command -v mysql >/dev/null 2>&1; then
  echo "mysql client not found. Install it or run the SQL file via another method."
  exit 1
fi

if [ -z "$DB_PASS" ]; then
  MYSQL_AUTH_ARGS=( -u"$DB_USER" -h"$DB_HOST" -P"$DB_PORT" )
else
  MYSQL_AUTH_ARGS=( -u"$DB_USER" -p"$DB_PASS" -h"$DB_HOST" -P"$DB_PORT" )
fi

echo "Applying ALTER/NOT NULL migration: $SQL_FILE"
mysql "${MYSQL_AUTH_ARGS[@]}" < "$SQL_FILE"

echo "Done. Tables updated to use defaults and NOT NULL where applicable."
