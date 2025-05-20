package main

import "net/http"

func (app *application) indexHandler(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	data.Articles = app.models.Articles.Latest()

	app.render(w, r, http.StatusOK, "index.tmpl", data)
}

func (app *application) viewArticleHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	article, err := app.models.Articles.GetBySlug(slug)
	if err != nil {
		app.notFoundErrorHandler(w, r)
		return
	}

	data := app.newTemplateData(r)
	data.Article = article

	app.render(w, r, http.StatusOK, "article.tmpl", data)
}
