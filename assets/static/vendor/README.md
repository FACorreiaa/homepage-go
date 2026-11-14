# Vendored JavaScript

This site has no bundler, no `package.json`, and no Node in the build — Tailwind
runs as a standalone binary and everything here is committed by hand and
embedded into the Go binary via `assets/assets.go`.

So third-party JS lives here, pinned. Update by re-downloading the same paths at
a new version and recording it below. Do not add a CDN `<script src>` instead:
the production CSP is `script-src 'self'` (see `docs/deployment.md`), and any
external origin is blocked outright.

| Library | Version | Source | Global |
|---|---|---|---|
| three | r*(see file banner)* | three.js | ES module via importmap |
| 3d-force-graph | *(see file)* | — | ES module |
| Alpine | *(see file)* | — | `window.Alpine` |
| htmx | *(see file)* | — | `window.htmx` |

## Notes

