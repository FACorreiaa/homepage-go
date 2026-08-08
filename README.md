# FC Software Studio Go

Go/templ implementation of the FC Software Studio site.

## Run

```sh
cp .env.example .env
go mod tidy
task dev
```

Open `http://localhost:7331` for templ live preview or `http://localhost:8090` for the app.

## Deployment

Production runs on k3s, reconciled by ArgoCD from
[`LuminaVault/LuminaVaultInfra`](https://github.com/LuminaVault/LuminaVaultInfra).
Merging to `main` is the whole deploy: CI builds the image, pushes it to GHCR,
and commits the new tag into the infra repo.

**See [`docs/deployment.md`](docs/deployment.md)** for deploying a code change,
adding env vars and secrets, changing resources, rolling back, and
troubleshooting.

---

# Legacy — decommissioned VPS setup

> [!WARNING]
> **Everything below this line is history, kept for reference only.**
>
> The single-VPS Docker Compose deployment was retired in `173c83a` when
> production moved to k3s + ArgoCD. There is no VPS to SSH into, the SSH deploy
> workflow is gone, and the `Caddyfile` is no longer read by anything — Traefik
> terminates TLS and applies the security headers now.
>
> `docker compose` is still useful for running the full stack locally. The
> `scripts/` referenced below and the vault-sync systemd units in `deploy/` are
> not part of the current deployment path.

## VPS Deployment (legacy)

### 1) Clone on VPS

```sh
cd /opt
git clone git@github.com:FACorreiaa/homepage-go.git facorreia-site-go
cd facorreia-site-go
```

If SSH access fails, copy/add a repo-specific key to GitHub and use:

```sh
GIT_SSH_COMMAND='ssh -i /path/to/key -o IdentitiesOnly=yes' \
  git clone git@github.com:FACorreiaa/homepage-go.git facorreia-site-go
```

### 2) Create required runtime folders

```sh
mkdir -p /opt/facorreia-site-go/vault/raw
```

`vault/raw` is the only required mount layout for bookmarks. The app reads
content from `/app/vault` inside the container.

### 3) Configure environment

```sh
cd /opt/facorreia-site-go
cp .env.example .env
```

Set at least these values in `.env`:

- `GO_ENV=production`
- `SESSION_SECRET` (strong random value)
- `ADMIN_EMAIL`
- `ADMIN_PASSWORD`
- `HOST_VAULT_PATH=/opt/facorreia-site-go/vault`

Optional:

- `DISCORD_WEBHOOK_URL`
- `CALENDLY_URL`

The app also uses these internal paths from `docker-compose.yml`:

- `DATABASE_PATH=/var/lib/facorreia-site/studio.sqlite`
- `BLOG_STATS_PATH=/var/lib/facorreia-site/blog_stats.json`
- `VAULT_PATH=/app/vault`

Protect `.env`:

```sh
chmod 600 .env
```

### 4) Build and run

```sh
git pull
docker compose build
docker compose up -d
docker compose ps
```

`studio_data` is the persistent Docker volume used for SQLite and blog stats. You
do not need a host folder at `/var/lib/facorreia-site`.

## Production (legacy)

The production compose file runs the Go app on port `8090`, stores SQLite and
blog analytics in a named Docker volume, and mounts the vault read-only.

Override these for production:

- `SESSION_SECRET`
- `ADMIN_EMAIL`
- `ADMIN_PASSWORD`

Optional environment:

- `DISCORD_WEBHOOK_URL`
- `CALENDLY_URL`
- `HOST_VAULT_PATH`, the host directory that contains `raw/`

Build and run:

```sh
docker compose build
docker compose up -d
```

Then open `http://localhost:8090`.

Health check endpoint:

```sh
curl -fsS http://localhost:8090/healthz
```

The container reads:

- `DATABASE_PATH=/var/lib/facorreia-site/studio.sqlite`
- `BLOG_STATS_PATH=/var/lib/facorreia-site/blog_stats.json`
- `VAULT_PATH=/app/vault`

For fresh production bookmark content, sync the Obsidian vault to
`HOST_VAULT_PATH`. The app refreshes the in-memory bookmark index in the
background when it becomes stale, using `BOOKMARK_INDEX_TTL` if set or 5
minutes by default. No container restart is required for normal syncs.

On the VPS, verify the vault sync timer with:

```sh
systemctl list-timers facorreia-vault-sync.timer --no-pager
journalctl -u facorreia-vault-sync -n 50 --no-pager
find /opt/facorreia-site-go/vault/raw -type f -name '*.md' | wc -l
```

## Simple Personal-Site Infra (legacy)

Keep the production setup intentionally small:

- Caddy terminates HTTPS and proxies to the Go container.
- Docker Compose runs one app service with `restart: unless-stopped`.
- SQLite and blog stats live in the `studio_data` Docker volume.
- Proposal leads are saved in SQLite first; Discord notification is best-effort.
- External uptime monitoring should check `https://facorreia.com/healthz`.

Back up the persistent data from the running container:

```sh
./scripts/backup-site-data.sh
```

Set `BACKUP_DIR=/path/to/backups` to write outside the repo. Schedule that
script with cron or a systemd timer on the VPS. For this site, that is enough
operational machinery unless traffic or lead volume changes materially.

## Cutover Checks (legacy)

Validate a copied Swift SQLite database before switching traffic:

```sh
./scripts/validate-cutover-db.sh ../facorreia-site/studio.sqlite
```

Then run:

```sh
go test ./...
go build -o /tmp/facorreia-site-go-check ./main.go
docker compose build
```

Smoke-test:

```sh
curl -I https://www.facorreia.com
curl -I https://facorreia.com
curl -I https://facorreia.com/assets/css/output.css
curl -I https://facorreia.com/assets/static/sw.js
curl -I https://facorreia.com/assets/static/manifest.json
curl -I https://facorreia.com/assets/static/vendor/htmx/htmx.min.js
```

Also verify:

- `/?` and public routes load
- `/proposal` POST works against production SQLite
- `/admin/login` and `/admin/dashboard` behavior
- `/bookmarks` renders after content appears in `vault/raw` and the bookmark index TTL expires
- service worker/caching headers for static assets are present
