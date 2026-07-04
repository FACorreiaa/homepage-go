package layouts

import (
	"context"
	"strings"
)

type ctxKey int

const requestPathKey ctxKey = 0

const siteBaseURL = "https://www.facorreia.com"

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
