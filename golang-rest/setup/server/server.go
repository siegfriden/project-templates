package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

func ListenAndServe(handler http.Handler, port int, config *Config) error {
	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		Handler:      handler,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	}

	return runServerWithGracefulShutdown(server, config.ShutdownTimeout)
}

func runServerWithGracefulShutdown(server *http.Server, shutdownTimeout time.Duration) error {
	// Start the server in a separate goroutine.
	serverErr := make(chan error, 1)
	go func() {
		serverErr <- server.ListenAndServe()
	}()

	// Set up termination signal notification.
	signalCtx, signalCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer signalCancel()

	// Wait for server error or termination signal.
	select {
	case err := <-serverErr:
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server failed to start: %w", err)
		}
		return nil
	case <-signalCtx.Done():
		if shutdownTimeout <= 0 {
			return server.Close() // No graceful shutdown, exit immediately.
		}

		// Graceful shutdown with timeout.
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("graceful shutdown failed: %w", err)
		}
		return nil
	}
}
