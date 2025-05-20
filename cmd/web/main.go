package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type application struct {
	debug         bool
	logger        *slog.Logger
	templateCache map[string]*template.Template
}

func main() {
	loadEnv()

	port := os.Getenv("PORT")
	debug := strings.ToLower(os.Getenv("DEBUG")) == "true"

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		debug:         debug,
		logger:        logger,
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: app.routes(),
	}

	logger.Info("Starting server on port " + port)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(0)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Unable to load env")
	}
}
