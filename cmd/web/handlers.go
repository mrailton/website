package main

import "net/http"

func (app *application) indexHandler(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	data.Articles = app.models.Articles.Latest()

	app.render(w, r, http.StatusOK, "index.tmpl", data)
}
