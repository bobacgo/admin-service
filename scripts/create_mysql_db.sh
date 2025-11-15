#!/usr/bin/env zsh
# Script to apply MySQL initialization SQL for admin-service.
# Usage:
#   DB_USER=root DB_PASS=pass DB_HOST=127.0.0.1 DB_PORT=3306 DB_NAME=admin_db ./scripts/create_mysql_db.sh

set -euo pipefail

DB_USER=${DB_USER:-root}
DB_PASS=${DB_PASS:-admin123}
DB_HOST=${DB_HOST:-127.0.0.1}
DB_PORT=${DB_PORT:-3306}
DB_NAME=${DB_NAME:-admin_db}
SQL_FILE="$(cd "$(dirname "$0")/.." && pwd)/migrations/mysql/init_mysql.sql"

if ! command -v mysql >/dev/null 2>&1; then
  echo "mysql client not found. Install it or run the SQL file via another method."
  exit 1
fi

# Build mysql CLI args
if [ -z "$DB_PASS" ]; then
  MYSQL_AUTH_ARGS=( -u"$DB_USER" -h"$DB_HOST" -P"$DB_PORT" )
else
  MYSQL_AUTH_ARGS=( -u"$DB_USER" -p"$DB_PASS" -h"$DB_HOST" -P"$DB_PORT" )
fi

echo "Applying SQL file: $SQL_FILE"
mysql "${MYSQL_AUTH_ARGS[@]}" < "$SQL_FILE"

echo "Done. Database '$DB_NAME' and tables should be created (if permissions allowed)."
