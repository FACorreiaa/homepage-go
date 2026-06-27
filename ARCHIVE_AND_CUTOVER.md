# Archive Swift and Cut Over to Go

This runbook preserves the existing Swift/Vapor project before making this Go/templ project the active `FACorreiaa/homepage` repository and `www.facorreia.com` deployment.

## 1. Validate Data Before Traffic Moves

Run the cutover database validation from this Go repo against the copied Swift SQLite database:

```sh
./scripts/validate-cutover-db.sh ../facorreia-site/studio.sqlite
```

Do not continue until the script reports success and the expected proposal, blog, and user tables are present.

## 2. Archive the Swift Repository

In the current Swift/Vapor repo:

```sh
git status --short
git push origin main
git tag swift-vapor-final-2026-06-27
git push origin swift-vapor-final-2026-06-27
```

Then in GitHub:

1. Rename `FACorreiaa/homepage` to `FACorreiaa/homepage-swift-archive`.
2. Confirm all branches and tags are visible after the rename.
3. Mark `FACorreiaa/homepage-swift-archive` as archived.

The Swift repo must remain recoverable. Do not overwrite it in place.

## 3. Prepare the Go Repository Remote

After the Swift rename frees the `homepage` name, create a new empty GitHub repository named `FACorreiaa/homepage`.

In this Go repo:

```sh
git remote set-url origin git@github.com:FACorreiaa/homepage.git
git branch -M main
git status --short
git add .
git commit -m "Cut over homepage to Go templ site"
git push -u origin main
```

Confirm GitHub shows this Go/templ codebase as the active `FACorreiaa/homepage` project.

## 4. Preserve Runtime Data on the Host

The Go compose service stores runtime state in the `studio_data` Docker volume:

- SQLite: `/var/lib/facorreia-site/studio.sqlite`
- Blog stats: `/var/lib/facorreia-site/blog_stats.json`
- Vault mount: `${HOST_VAULT_PATH:-./vault}:/app/vault:ro`

If the Swift and Go services run on the same host, copy the Swift Docker volume data into the Go volume before starting the Go app. Adjust the old volume name if needed:

```sh
docker run --rm \
  -v homepage_swift_data:/from:ro \
  -v facorreia-site-go_studio_data:/to \
  alpine sh -c "cp -a /from/. /to/"
```

Then validate the copied SQLite file from inside the Go repo or host path that maps to the volume.

## 5. Build and Start the Go Site

```sh
go tool templ generate
tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css
go test ./...
go build -o /tmp/facorreia-site-go-check ./main.go
docker compose build
docker compose up -d
```

The app must remain bound to port `8090`.

## 6. Reload Caddy

Use the provided `Caddyfile` from this repo:

```caddy
facorreia.com {
	redir https://www.facorreia.com{uri} permanent
}

www.facorreia.com {
	encode zstd gzip
	reverse_proxy localhost:8090
}
```

Reload Caddy after copying the file into place:

```sh
caddy reload --config /etc/caddy/Caddyfile
```

## 7. Verify Production

Check the public site:

```sh
curl -I https://www.facorreia.com
curl -I https://facorreia.com
```

Expected behavior:

- `https://www.facorreia.com` serves the Go site.
- `https://facorreia.com` returns a permanent redirect to `https://www.facorreia.com`.
- `/assets/static/vendor/gsap/gsap.min.js` returns 200.
- `/assets/static/vendor/gsap/ScrollTrigger.min.js` returns 200.
- `/proposal` accepts a valid test submission into production SQLite.
- `/blog` shows seeded posts and preserves blog stats.
- `/bookmarks` shows mounted vault content after a container restart.
- `/admin/dashboard` redirects while logged out, login works, and logout clears access.
