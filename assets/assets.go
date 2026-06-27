package assets

import "embed"

//go:embed css/* fonts/* static
var Assets embed.FS
