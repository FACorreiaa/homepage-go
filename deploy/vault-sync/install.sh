#!/usr/bin/env bash
# Install the vault-sync systemd units on the homepage server.
# Run as root from the repo root: bash deploy/vault-sync/install.sh
set -euo pipefail

UNIT_DIR=/etc/systemd/system
SRC_DIR="$(cd "$(dirname "$0")" && pwd)"
KEY=/root/.ssh/id_ed25519_vault_sync

if [[ ! -f "$KEY" ]]; then
  echo "Generating dedicated sync key $KEY"
  ssh-keygen -t ed25519 -f "$KEY" -N '' -C vault-sync@homepage
  echo
  echo "Add this public key to root@78.46.192.73:~/.ssh/authorized_keys"
  echo "(ideally restricted: restrict,command=\"rrsync -ro /root/.hermes/obsidian-vault/FACorreia/raw/\"):"
  echo
  cat "$KEY.pub"
  echo
fi

mkdir -p /opt/facorreia-site-go/vault/raw

cp "$SRC_DIR/facorreia-vault-sync.service" "$UNIT_DIR/"
cp "$SRC_DIR/facorreia-vault-sync.timer" "$UNIT_DIR/"
systemctl daemon-reload
systemctl enable --now facorreia-vault-sync.timer

echo "Timer installed:"
systemctl list-timers facorreia-vault-sync.timer --no-pager || true
echo
echo "Run the first sync manually with:"
echo "  systemctl start facorreia-vault-sync.service && journalctl -u facorreia-vault-sync -n 20 --no-pager"
