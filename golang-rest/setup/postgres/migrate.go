package postgres

import (
	"context"
	"database/sql"
	"io/fs"
	"os"

	"github.com/pressly/goose/v3"
)

// MigrateDir applies migrations with Goose and returns the list of applied migration files.
func MigrateDir(db *sql.DB, dir string) ([]string, error) {
	return MigrateFS(db, os.DirFS(dir))
}

// MigrateFS applies migrations with Goose and returns the list of applied migration files.
func MigrateFS(db *sql.DB, fs fs.FS) ([]string, error) {
	p, err := goose.NewProvider("postgres", db, fs)
	if err != nil {
		return nil, err
	}
	results, err := p.Up(context.Background())
	if err != nil {
		return nil, err
	}

	// Get the paths of applied migrations.
	// Useful for logging or debugging.
	sources := make([]string, len(results))
	for i, result := range results {
		sources[i] = result.Source.Path
	}
	return sources, nil
}
