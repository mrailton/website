package main

import (
	"net/http"

	"markrailton.com/ui"
)

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.Handle("GET /static/", http.FileServerFS(ui.Files))
	router.HandleFunc("GET /", app.notFoundErrorHandler)

	router.HandleFunc("GET /{$}", app.indexHandler)
	router.HandleFunc("GET /blog/{slug}", app.viewArticleHandler)

	return router
}
