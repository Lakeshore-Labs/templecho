package assets

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed js/* css/*
var Assets embed.FS

// GetAssetsFS returns the assets subdirectory as an http.FileSystem
func GetAssetsFS() http.FileSystem {
	return http.FS(Assets)
}

// GetJSFS returns just the JS assets
func GetJSFS() http.FileSystem {
	sub, _ := fs.Sub(Assets, "js")
	return http.FS(sub)
}

// GetCSSFS returns just the CSS assets
func GetCSSFS() http.FileSystem {
	sub, _ := fs.Sub(Assets, "css")
	return http.FS(sub)
}