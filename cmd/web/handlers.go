package main

import "net/http"

func (app *application) indexHandler(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, r, http.StatusOK, "index.tmpl", data)
}
