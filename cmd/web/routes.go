package main

import (
	"net/http"

	"markrailton.com/ui"
)

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.Handle("GET /static/", http.FileServerFS(ui.Files))

	router.HandleFunc("GET /", app.indexHandler)

	return router
}
