package postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Config struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration

	WaitForReadyTimeout    time.Duration
	WaitForReadyMaxBackoff time.Duration
}

func Connect(connStr string, config *Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if config != nil {
		db.SetMaxOpenConns(config.MaxOpenConns)
		db.SetMaxIdleConns(config.MaxIdleConns)
		db.SetConnMaxLifetime(config.ConnMaxLifetime)
	}

	if err := waitForReady(
		db,
		config.WaitForReadyTimeout,
		config.WaitForReadyMaxBackoff,
	); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func waitForReady(db *sql.DB, timeout time.Duration, maxBackoff time.Duration) error {
	// Backoff duration can't be less than 1 second.
	if maxBackoff < time.Second {
		maxBackoff = time.Second
	}

	start := time.Now()
	sleep := time.Second

	// Retry database ping until success or timeout.
	for {
		err := db.Ping()
		if err == nil {
			return nil
		}
		if time.Since(start) > timeout {
			return fmt.Errorf("ping database: %w", err)
		}

		if sleep > maxBackoff {
			sleep = maxBackoff
		}
		time.Sleep(sleep)
		sleep *= 2 // Exponential backoff.
	}
}
