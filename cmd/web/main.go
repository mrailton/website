package main

import (
	"context"
	"database/sql"
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
	_ "github.com/lib/pq"

	"markrailton.com/internal/data"
)

type config struct {
	port  string
	env   string
	debug bool
	db    struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}

type application struct {
	config        config
	logger        *slog.Logger
	models        data.Models
	templateCache map[string]*template.Template
	wg            sync.WaitGroup
	db            *sql.DB
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

	cfg := config{
		port:  os.Getenv("PORT"),
		debug: os.Getenv("DEBUG") == "true",
		env:   os.Getenv("APP_ENV"),
	}

	cfg.db.dsn = os.Getenv("DSN")
	cfg.db.maxOpenConns = 25
	cfg.db.maxIdleConns = 25
	cfg.db.maxIdleTime = 15 * time.Minute

	return cfg
}

func buildApplication() *application {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	templateCache := loadTemplateCache(logger)
	cfg := buildConfig()

	db := openDB(cfg, logger)

	return &application{
		config:        cfg,
		logger:        logger,
		templateCache: templateCache,
		models:        data.NewModels(db),
		db:            db,
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

		app.logger.Info("closing database connection")
		if err := app.db.Close(); err != nil {
			app.logger.Error("error closing database connection", "error", err)
		}

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

func openDB(cfg config, logger *slog.Logger) *sql.DB {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()

		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("database connection established")
	return db
}
