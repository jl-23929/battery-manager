package main

import (
	"net/http"

	"github.com/jl-23929/battery-manager/assets"
)

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(assets.EmbeddedFiles))

	mux.HandleFunc("GET /{$}", homepage)
	mux.HandleFunc("GET /newbattery", newBattery)
	mux.HandleFunc("GET /viewbattery/{id}", viewBattery)
	mux.HandleFunc("GET /generate", generateCode)

	mux.HandleFunc("GET /health", health)

	return mux
}
