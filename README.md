# FC Software Studio Go

Go/templ implementation of the FC Software Studio site.

## Run

```sh
cp .env.example .env
go mod tidy
task dev
```

Open `http://localhost:7331` for templ live preview or `http://localhost:8090` for the app.

## VPS Deployment

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

## Production

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

The container reads:

- `DATABASE_PATH=/var/lib/facorreia-site/studio.sqlite`
- `BLOG_STATS_PATH=/var/lib/facorreia-site/blog_stats.json`
- `VAULT_PATH=/app/vault`

For fresh production bookmark content, sync the Obsidian vault to
`HOST_VAULT_PATH` and restart the container so the in-memory bookmark index is
rebuilt.

## Cutover Checks

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
curl -I https://www.facorreia.com/assets/static/vendor/gsap/gsap.min.js
curl -I https://www.facorreia.com/assets/static/vendor/gsap/ScrollTrigger.min.js
```

Also verify:

- `/?` and public routes load
- `/proposal` POST works against production SQLite
- `/admin/login` and `/admin/dashboard` behavior
- `/bookmarks` renders after content in `vault/raw` and container restart
- service worker/caching headers for static assets are present
