// Package htmx provides embedded HTMX JavaScript library
package htmx

import (
	_ "embed"
	"net/http"
)

// Version is the HTMX library version
const Version = "1.9.10"

//go:embed htmx.min.js
var htmxJS []byte

// JavaScript returns the embedded HTMX JavaScript content
func JavaScript() []byte {
	return htmxJS
}

// Handler returns an http.Handler that serves the HTMX JavaScript
func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Header().Set("Cache-Control", "public, max-age=31536000") // Cache for 1 year
		w.Write(htmxJS)
	})
}

// ScriptTag returns an HTML script tag for including HTMX
func ScriptTag(path string) string {
	return `<script src="` + path + `"></script>`
}