package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"gitlab.com/siegfriden/project-templates/golang-rest/setup/env"
	"gitlab.com/siegfriden/project-templates/golang-rest/setup/postgres"
	"gitlab.com/siegfriden/project-templates/golang-rest/setup/server"
)

func main() {
	logger := setupLogger()
	logger.Info("Starting Go server...")

	config := loadConfig(logger)
	db := connectPostgres(logger, config.postgresURL)
	defer db.Close()

	handler := setupAPIHandler(logger, db)
	startServer(logger, handler, config.serverPort)
}

func setupLogger() *slog.Logger {
	// Hostname will be useful for pod/instance identification in distributed logging.
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	// Create structured JSON logger.
	return slog.New(slog.NewJSONHandler(os.Stderr, nil).WithAttrs([]slog.Attr{
		slog.String("service", "go-server"),
		slog.String("hostname", hostname),
	}))
}

type appConfig struct {
	postgresURL string
	serverPort  int
}

func loadConfig(logger *slog.Logger) *appConfig {
	// Try to load .env file.
	if err := godotenv.Load(); err != nil {
		// It's okay if .env doesn't exist.
		logger.Info("No .env file found.")
	}

	// Load environment variables into the config struct.
	// Uses custom `env` package for type-conversion and validation.
	envLoader := env.NewLoader()
	config := &appConfig{
		postgresURL: envLoader.Required("POSTGRES_URL"),
		serverPort:  envLoader.OptionalInt("SERVER_PORT", 8080),
	}

	// Missing or mistyped environment variables will cause an error here.
	if err := envLoader.Validate(); err != nil {
		logger.Error("Failed to load configuration.", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return config
}

func connectPostgres(logger *slog.Logger, postgresURL string) *sql.DB {
	db, err := postgres.Connect(postgresURL, &postgres.Config{
		MaxOpenConns:    50,
		MaxIdleConns:    10,
		ConnMaxLifetime: 15 * time.Minute,

		// Retry connecting for 5 seconds in case app starts before DB (e.g., Docker Compose).
		WaitForReadyTimeout:    5 * time.Second,
		WaitForReadyMaxBackoff: 1 * time.Second,
	})
	if err != nil {
		logger.Error("Failed to connect to PostgreSQL database.", slog.String("error", err.Error()))
		os.Exit(1)
	}
	return db
}

func setupAPIHandler(logger *slog.Logger, db *sql.DB) http.Handler {
	router := chi.NewRouter()
	// TODO: Add middlewares (logging, CORS, auth, etc.)

	router.Route("/v1", func(r chi.Router) {
		// Public endpoints...
		// r.HandleFunc("/register", ...)
		// r.HandleFunc("/login", ...)
		// r.HandleFunc("/logout", ...)
		// r.HandleFunc("/verify-session", ...)

		// Protected endpoints...
		r.Group(func(r chi.Router) {
			// r.Use(auth.Require)
			// ...
		})
	})

	return router
}

func startServer(logger *slog.Logger, router http.Handler, port int) {
	logger.Info("Starting HTTP server.", slog.Int("port", port))

	// Start HTTP server with graceful shutdown.
	err := server.ListenAndServe(router, port, &server.Config{
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
		IdleTimeout:     60 * time.Second,
		ShutdownTimeout: 10 * time.Second,
	})
	if err != nil {
		logger.Error("HTTP server failed.", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
