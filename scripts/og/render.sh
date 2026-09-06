#!/usr/bin/env bash
# Render the 1200×630 Open Graph PNG from scripts/og/facorreia-og.html.
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
HTML="$ROOT/scripts/og/facorreia-og.html"
OUT="$ROOT/assets/static/og.png"
CHROME="${CHROME:-/Applications/Google Chrome.app/Contents/MacOS/Google Chrome}"
if [[ ! -x "$CHROME" ]]; then
  echo "Chrome not found at $CHROME" >&2
  exit 1
fi
"$CHROME" \
  --headless=new \
  --disable-gpu \
  --hide-scrollbars \
  --force-device-scale-factor=1 \
  --window-size=1200,630 \
  --virtual-time-budget=4000 \
  --screenshot="$OUT" \
  "file://$HTML"
sips -z 630 1200 "$OUT" >/dev/null
sips -g pixelWidth -g pixelHeight -g format "$OUT"
