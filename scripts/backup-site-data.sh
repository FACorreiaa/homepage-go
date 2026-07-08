#!/usr/bin/env sh
# Create a timestamped backup of the site's persistent runtime data.
# Run from the directory that contains docker-compose.yml.
set -eu

SERVICE="${SERVICE:-app}"
BACKUP_DIR="${BACKUP_DIR:-./backups}"
STAMP="$(date -u +%Y%m%dT%H%M%SZ)"
OUT="$BACKUP_DIR/facorreia-site-data-$STAMP.tar.gz"

mkdir -p "$BACKUP_DIR"

docker compose exec -T "$SERVICE" sh -c '
  set -eu
  cd /var/lib/facorreia-site
  tar -czf - studio.sqlite blog_stats.json 2>/dev/null
' > "$OUT"

echo "backup written: $OUT"
