package main

import (
	"context"
	"errors"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

type config struct {
	port  string
	env   string
	debug bool
}

type application struct {
	config        config
	logger        *slog.Logger
	templateCache map[string]*template.Template
	wg            sync.WaitGroup
}

func main() {
	app := buildApplication()
	srv := buildServer(app)

	runServer(app, srv)
}

func buildConfig() config {
	err := godotenv.Load()
	if err != nil {
		panic("Unable to load env")
	}

	return config{
		port:  os.Getenv("PORT"),
		debug: os.Getenv("DEBUG") == "true",
	}
}

func buildApplication() *application {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	templateCache := loadTemplateCache(logger)

	return &application{
		config:        buildConfig(),
		logger:        logger,
		templateCache: templateCache,
	}
}

func buildServer(app *application) *http.Server {
	return &http.Server{
		Addr:    ":" + app.config.port,
		Handler: app.routes(),
	}
}

func runServer(app *application, srv *http.Server) {
	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		app.logger.Info("shutting down server", "signal", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		app.logger.Info("completing background tasks", "address", srv.Addr)

		app.wg.Wait()
		shutdownError <- nil
	}()

	app.logger.Info("starting server", "addr", srv.Addr, "env", app.config.env)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return
	}

	err = <-shutdownError
	if err != nil {
		return
	}

	app.logger.Info("stopped server", "addr", srv.Addr)
	os.Exit(0)
}
