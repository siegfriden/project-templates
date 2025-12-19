.PHONY: help build run test test-coverage clean deps migrate-up migrate-down migrate-status migrate-create install-tools

# Load variables from ".env" file if exists.
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

help: ## Show this help message.
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application.
	@echo "Building application..."
	go build -o bin/project-tracker ./main.go

run: ## Run the application.
	@echo "Running application..."
	go run main.go

test: ## Run tests.
	@echo "Running tests..."
	go test -v ./...

test-coverage: ## Run tests with coverage.
	@echo "Running tests with coverage..."
	go test -v -race -buildvcs -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

clean: ## Clean build artifacts.
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -f coverage.out coverage.html

deps: ## Install dependencies.
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

migrate-up: ## Run database migrations up.
	@echo "Running migrations up..."
	goose -dir database/migrations postgres "$(POSTGRES_URL)" up

migrate-down: ## Run database migrations down.
	@echo "Running migrations down..."
	goose -dir database/migrations postgres "$(POSTGRES_URL)" down

migrate-reset: ## Reset database migrations.
	@echo "Resetting migrations..."
	goose -dir database/migrations postgres "$(POSTGRES_URL)" reset

migrate-status: ## Check database migration status.
	@echo "Checking migration status..."
	goose -dir database/migrations postgres "$(POSTGRES_URL)" status

migrate-create: ## Create a new migration (usage: make migrate-create name=migration_name).
	@echo "Creating migration: $(name)"
	goose -dir resources/migrations create $(name) sql

sqlc: ## Generate Go code from SQL schema and queries.
	@echo "Generating Go code from SQL..."
	sqlc generate

install-tools: ## Install development tools.
	@echo "Installing development tools..."
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
