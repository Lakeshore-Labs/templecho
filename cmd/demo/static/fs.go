package static

import (
	"embed"
	"io/fs"
)

//go:embed css/*.css js/*.js fonts/geist/*.woff2
var embeddedFS embed.FS

// FS returns the embedded filesystem
func FS() fs.FS {
	return embeddedFS
}