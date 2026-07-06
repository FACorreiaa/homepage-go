#!/usr/bin/env sh
# Drop external access to the app's published Docker port. Docker publishes
# :8090 on all interfaces and bypasses UFW, so this rule in the DOCKER-USER
# chain (which Docker always honours and never flushes) is what actually
# keeps the port private. Caddy reaches the app via the host loopback, which
# does not traverse DOCKER-USER, so the reverse proxy is unaffected.
#
# Idempotent: delete any existing copy, then insert. Re-run safely.
set -eu

IFACE="${EXT_IFACE:-eth0}"
PORT="${APP_PORT:-8090}"

for ipt in iptables ip6tables; do
  command -v "$ipt" >/dev/null 2>&1 || continue
  while "$ipt" -C DOCKER-USER -i "$IFACE" -p tcp --dport "$PORT" -j DROP 2>/dev/null; do
    "$ipt" -D DOCKER-USER -i "$IFACE" -p tcp --dport "$PORT" -j DROP
  done
  "$ipt" -I DOCKER-USER -i "$IFACE" -p tcp --dport "$PORT" -j DROP
done

echo "docker-port-firewall: dropping $IFACE tcp/$PORT from external"
