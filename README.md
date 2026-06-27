# FC Software Studio Go

Go/templ implementation of the FC Software Studio site.

## Run

```sh
cp .env.example .env
go mod tidy
task dev
```

Open `http://localhost:7331` for templ live preview or `http://localhost:8090` for the app.

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

Smoke-test `/`, `/projects`, `/projects/luminavault`, `/blog`,
`/blog/{slug}`, `/bookmarks`, `/api/graph`, `/proposal`, `/admin/login`, and
`/admin/dashboard`.
# homepage-go
# homepage-go
