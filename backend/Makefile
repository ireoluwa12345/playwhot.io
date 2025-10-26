BINARY_NAME = playwhot
CMD_PATH = ./cmd/web
MIGRATIONS_DIR = MIGRATIONS_DIR

DB_DSN ?= $(DB_DSN)

.PHONY: help build run install-goose migrate-up migrate-down migrate-status tidy test clean

help:
	@echo "Available targets:"
	@echo "  make build            -> compile the web binary"
	@echo "  make run              -> run the server using 'go run'"
	@echo "  make install-goose    -> install the goose CLI to \\$(GOBIN or GOPATH/bin)"
	@echo "  make migrate-up       -> apply all up migrations (requires DB_DSN)"
	@echo "  make migrate-down     -> apply down (rollback) one migration (requires DB_DSN)"
	@echo "  make migrate-status   -> show goose migration status (requires DB_DSN)"
	@echo "  make tidy             -> run 'go mod tidy'"
	@echo "  make test             -> run tests"
	@echo "  make clean            -> remove built binary"

build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) $(CMD_PATH)

run:
	@echo "Running server (use DB_DSN env to enable migrations at startup)..."
	go run $(CMD_PATH)

install-goose:
	@echo "Installing goose CLI..."
	go install github.com/pressly/goose/v3/cmd/goose@latest

%_require_dsn:
	@if [ -z "$(DB_DSN)" ]; then \
		echo "DB_DSN must be set. Example:"; \
		echo "  make migrate-up DB_DSN='user:pass@tcp(localhost:3306)/dbname?parseTime=true'"; \
		exit 1; \
	fi

migrate-up: %_require_dsn
	@echo "Applying up migrations from '$(MIGRATIONS_DIR)'..."
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_DSN)" up

migrate-down: %_require_dsn
	@echo "Rolling back last migration in '$(MIGRATIONS_DIR)'..."
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_DSN)" down

migrate-status: %_require_dsn
	@echo "Migration status for '$(MIGRATIONS_DIR)'..."
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_DSN)" status

tidy:
	go mod tidy

test:
	go test ./...

clean:
	@echo "Removing binary $(BINARY_NAME)"
	-rm -f $(BINARY_NAME)

