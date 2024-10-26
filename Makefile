.PHONY: dev build run
dev:
	@echo "Running bot..."
	@go run cmd/bot/main.go
build:
	@echo "Building bot..."
	@go build -o bin/bot cmd/bot/main.go
	@cp .env ./bin/.env
run: build
	@echo "Running bot..."
	@./bin/bot
