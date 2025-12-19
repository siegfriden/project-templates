package migrations

import "embed"

// EmbeddedFS contains all SQL migration files embedded at compile time.
// This allows the application binary to include migration files without requiring
// external file dependencies during deployment.
//
//go:embed *.sql
var EmbeddedFS embed.FS
