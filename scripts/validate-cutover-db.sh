#!/usr/bin/env sh
set -eu

if [ "$#" -ne 1 ]; then
  echo "usage: $0 /path/to/swift/studio.sqlite" >&2
  exit 2
fi

source_db="$1"
if [ ! -f "$source_db" ]; then
  echo "database not found: $source_db" >&2
  exit 2
fi

tmp_db="$(mktemp "${TMPDIR:-/tmp}/facorreia-cutover.XXXXXX.sqlite")"
cleanup() {
  rm -f "$tmp_db"
}
trap cleanup EXIT

cp "$source_db" "$tmp_db"

integrity="$(sqlite3 "$tmp_db" "PRAGMA integrity_check;")"
if [ "$integrity" != "ok" ]; then
  echo "sqlite integrity_check failed: $integrity" >&2
  exit 1
fi

sqlite3 "$tmp_db" < internal/db/schema.sql
sqlite3 "$tmp_db" "UPDATE proposal_leads SET created_at = datetime('now') WHERE created_at IS NULL OR created_at = '';"

for table in client_users proposal_leads blog_posts; do
  exists="$(sqlite3 "$tmp_db" "SELECT COUNT(*) FROM sqlite_master WHERE type = 'table' AND name = '$table';")"
  if [ "$exists" != "1" ]; then
    echo "missing required table: $table" >&2
    exit 1
  fi
done

sqlite3 "$tmp_db" "
BEGIN;
INSERT INTO proposal_leads (
  id, name, email, company, project_type, budget, timeline, details, created_at
) VALUES (
  '00000000-0000-0000-0000-000000000000',
  'Cutover Check',
  'cutover@example.com',
  NULL,
  'Website',
  NULL,
  NULL,
  'Compatibility insert',
  datetime('now')
);
INSERT INTO blog_posts (
  slug, title, summary, category, date, body
) VALUES (
  'cutover-check',
  'Cutover Check',
  'Compatibility insert',
  'System',
  date('now'),
  'ok'
);
ROLLBACK;
"

echo "cutover database validation passed for $source_db"
echo "client_users: $(sqlite3 "$tmp_db" "SELECT COUNT(*) FROM client_users;")"
echo "proposal_leads: $(sqlite3 "$tmp_db" "SELECT COUNT(*) FROM proposal_leads;")"
