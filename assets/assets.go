package assets

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"strings"
	"sync"
)

//go:embed css/* fonts/* static
var Assets embed.FS

// Immutable marks the asset trees whose contents can never change under a given
// path. Vendored libraries carry their version in the path and fonts are
// content-stable, so those are safe to pin for a year. Everything else this site
// serves — its own CSS and JS — is rebuilt on every deploy at a fixed path, and
// pinning those is how a returning visitor ends up running new HTML against a
// year-old stylesheet.
var immutablePrefixes = []string{
	"/assets/static/vendor/",
	"/assets/fonts/",
}

// IsImmutable reports whether a request path may be cached indefinitely.
func IsImmutable(requestPath string) bool {
	for _, prefix := range immutablePrefixes {
		if strings.HasPrefix(requestPath, prefix) {
			return true
		}
	}
	return false
}

var versions sync.Map

// Version is a short content hash of an embedded asset, for use as a cache
// buster in the URL. Revalidation headers alone cannot reach a browser that
// already holds a copy pinned as immutable — only a URL it has never seen can.
func Version(name string) string {
	if v, ok := versions.Load(name); ok {
		return v.(string)
	}
	body, err := Assets.ReadFile(name)
	if err != nil {
		return "0"
	}
	sum := sha256.Sum256(body)
	v := hex.EncodeToString(sum[:])[:10]
	versions.Store(name, v)
	return v
}

// URL returns the public path for an embedded asset, fingerprinted so that a
// change to its contents changes its URL.
func URL(name string) string {
	return "/assets/" + name + "?v=" + Version(name)
}
