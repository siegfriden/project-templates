package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Loader provides methods to load and convert environment variables.
// It keeps track of missing and malformed variables.
// After loading, call Validate() to check for errors.
type Loader struct {
	missing   []string // Mandatory variables that are missing.
	malformed []string // Variables that failed type parsing.
}

func NewLoader() *Loader {
	return &Loader{}
}

// Validate returns an error if any required environment variables are missing
// or if any variables failed type parsing.
func (l *Loader) Validate() error {
	var errors []string

	if len(l.missing) > 0 {
		errors = append(errors, fmt.Sprintf("missing required env: %s", strings.Join(l.missing, ", ")))
	}

	if len(l.malformed) > 0 {
		errors = append(errors, fmt.Sprintf("malformed env: %s", strings.Join(l.malformed, ", ")))
	}

	if len(errors) > 0 {
		return fmt.Errorf("%s", strings.Join(errors, "; "))
	}

	return nil
}

func (l *Loader) Required(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		l.missing = append(l.missing, key)
	}
	return value
}

func (l *Loader) Optional(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func (l *Loader) RequiredInt(key string) int {
	valueStr, ok := os.LookupEnv(key)
	if !ok {
		l.missing = append(l.missing, key)
		return 0
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		l.malformed = append(l.malformed, key)
		return 0
	}
	return value
}

func (l *Loader) OptionalInt(key string, fallback int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return fallback
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		l.malformed = append(l.malformed, key)
		return fallback
	}
	return value
}

func (l *Loader) RequiredBool(key string) bool {
	valueStr, ok := os.LookupEnv(key)
	if !ok {
		l.missing = append(l.missing, key)
		return false
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		l.malformed = append(l.malformed, key)
		return false
	}
	return value
}

func (l *Loader) OptionalBool(key string, fallback bool) bool {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return fallback
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		l.malformed = append(l.malformed, key)
		return fallback
	}
	return value
}

func (l *Loader) RequiredDuration(key string) time.Duration {
	valueStr, ok := os.LookupEnv(key)
	if !ok {
		l.missing = append(l.missing, key)
		return 0
	}
	value, err := time.ParseDuration(valueStr)
	if err != nil {
		l.malformed = append(l.malformed, key)
		return 0
	}
	return value
}

func (l *Loader) OptionalDuration(key string, fallback time.Duration) time.Duration {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return fallback
	}
	value, err := time.ParseDuration(valueStr)
	if err != nil {
		l.malformed = append(l.malformed, key)
		return fallback
	}
	return value
}
