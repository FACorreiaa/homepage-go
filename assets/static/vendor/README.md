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
| GSAP | 3.15.0 | `cdn.jsdelivr.net/npm/gsap@3.15.0/dist/gsap.min.js` | `window.gsap` |
| ScrollTrigger | 3.15.0 | `cdn.jsdelivr.net/npm/gsap@3.15.0/dist/ScrollTrigger.min.js` | `window.ScrollTrigger` |
| Lenis | 1.3.26 | `cdn.jsdelivr.net/npm/lenis@1.3.26/dist/lenis.min.js` | `globalThis.Lenis` |
| Alpine | *(see file)* | — | `window.Alpine` |
| htmx | *(see file)* | — | `window.htmx` |

## Notes

- GSAP and ScrollTrigger are UMD builds and must load as plain
  `<script defer src="...">`, not as modules. ScrollTrigger calls
  `gsap.registerPlugin` itself on load, so gsap.min.js must come first.
- **GSAP ships no licence file on npm.** Its `package.json` declares
  *"Standard 'no charge' license: https://gsap.com/standard-license"*, and the
  full notice is in the `@license` banner at the top of each minified file. That
  banner is why these files are not re-minified or stripped.
- Lenis is MIT; `lenis/LICENSE` is the upstream file.
- The `//# sourceMappingURL` comment was removed from `lenis.min.js`. The `.map`
  file is not vendored, so leaving it would 404 in devtools on every page load.
