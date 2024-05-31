package assets

import (
	"embed"
)

//go:embed "emails" "migrations" "env"
var EmbeddedFiles embed.FS
