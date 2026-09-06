package layouts

import (
	"context"
	"strings"
)

type ctxKey int

const requestPathKey ctxKey = 0

const siteBaseURL = "https://facorreia.com"

// defaultOGImagePath is the 1200×630 share card. Square app icons are not
// Open Graph images — crawlers crop them into a landscape slot.
const defaultOGImagePath = "/assets/static/og.png"

func absoluteAssetURL(path string) string {
	if path == "" {
		return siteBaseURL + defaultOGImagePath
	}
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return path
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return siteBaseURL + path
}

// WithRequestPath stores the request path so BaseLayout can emit
// canonical/og:url tags without every page threading it through props.
func WithRequestPath(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, requestPathKey, path)
}

func canonicalURL(ctx context.Context) string {
	p, _ := ctx.Value(requestPathKey).(string)
	if p == "" || !strings.HasPrefix(p, "/") {
		p = "/"
	}
	if p != "/" {
		p = strings.TrimSuffix(p, "/")
	}
	return siteBaseURL + p
}
