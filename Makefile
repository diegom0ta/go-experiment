# Makefile for Go Experiment Project

# Variables
APP_NAME=go-experiment
GO_VERSION=1.25
DOCKER_COMPOSE=docker compose
DOCKER_IMAGE_NAME=$(APP_NAME)-app

# Colors for output
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[1;33m
BLUE=\033[0;34m
NC=\033[0m # No Color

.PHONY: help build run test clean docker-build docker-up docker-down docker-logs lint fmt vet deps migration-up migration-down db-reset

# Default target
help: ## Show this help message
	@echo "$(BLUE)Available commands:$(NC)"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  $(GREEN)%-20s$(NC) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Go commands
build: ## Build the application
	@echo "$(YELLOW)Building application...$(NC)"
	go build -o bin/$(APP_NAME) .
	@echo "$(GREEN)Build completed!$(NC)"

run: ## Run the application locally
	@echo "$(YELLOW)Running application...$(NC)"
	go run main.go

test: ## Run tests
	@echo "$(YELLOW)Running tests...$(NC)"
	go test -v ./...
	@echo "$(GREEN)Tests completed!$(NC)"

test-coverage: ## Run tests with coverage
	@echo "$(YELLOW)Running tests with coverage...$(NC)"
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

bench: ## Run benchmarks
	@echo "$(YELLOW)Running benchmarks...$(NC)"
	go test -bench=. -benchmem ./...

# Code quality
lint: ## Run golangci-lint
	@echo "$(YELLOW)Running linter...$(NC)"
	golangci-lint run
	@echo "$(GREEN)Linting completed!$(NC)"

fmt: ## Format Go code
	@echo "$(YELLOW)Formatting code...$(NC)"
	go fmt ./...
	@echo "$(GREEN)Code formatted!$(NC)"

vet: ## Run go vet
	@echo "$(YELLOW)Running go vet...$(NC)"
	go vet ./...
	@echo "$(GREEN)Vet completed!$(NC)"

# Dependencies
deps: ## Download and tidy dependencies
	@echo "$(YELLOW)Downloading dependencies...$(NC)"
	go mod download
	go mod tidy
	@echo "$(GREEN)Dependencies updated!$(NC)"

deps-upgrade: ## Upgrade all dependencies
	@echo "$(YELLOW)Upgrading dependencies...$(NC)"
	go get -u ./...
	go mod tidy
	@echo "$(GREEN)Dependencies upgraded!$(NC)"

# Docker commands
docker-build: ## Build Docker image
	@echo "$(YELLOW)Building Docker image...$(NC)"
	$(DOCKER_COMPOSE) build
	@echo "$(GREEN)Docker image built!$(NC)"

docker-up: ## Start all services with Docker Compose
	@echo "$(YELLOW)Starting services...$(NC)"
	$(DOCKER_COMPOSE) up -d
	@echo "$(GREEN)Services started!$(NC)"

docker-up-build: ## Build and start all services
	@echo "$(YELLOW)Building and starting services...$(NC)"
	$(DOCKER_COMPOSE) up -d --build
	@echo "$(GREEN)Services built and started!$(NC)"

docker-down: ## Stop all services
	@echo "$(YELLOW)Stopping services...$(NC)"
	$(DOCKER_COMPOSE) down
	@echo "$(GREEN)Services stopped!$(NC)"

docker-down-volumes: ## Stop all services and remove volumes
	@echo "$(YELLOW)Stopping services and removing volumes...$(NC)"
	$(DOCKER_COMPOSE) down -v
	@echo "$(GREEN)Services stopped and volumes removed!$(NC)"

docker-logs: ## Show logs for all services
	$(DOCKER_COMPOSE) logs -f

docker-logs-app: ## Show logs for the application
	$(DOCKER_COMPOSE) logs -f app

docker-logs-db: ## Show logs for the database
	$(DOCKER_COMPOSE) logs -f postgres

docker-logs-redis: ## Show logs for Redis
	$(DOCKER_COMPOSE) logs -f redis

docker-ps: ## Show running containers
	$(DOCKER_COMPOSE) ps

docker-restart: ## Restart all services
	@echo "$(YELLOW)Restarting services...$(NC)"
	$(DOCKER_COMPOSE) restart
	@echo "$(GREEN)Services restarted!$(NC)"

docker-restart-app: ## Restart only the application
	@echo "$(YELLOW)Restarting application...$(NC)"
	$(DOCKER_COMPOSE) restart app
	@echo "$(GREEN)Application restarted!$(NC)"

# Database commands
db-reset: ## Reset database (remove volume and restart)
	@echo "$(YELLOW)Resetting database...$(NC)"
	$(DOCKER_COMPOSE) down
	docker volume rm go-experiment_pgdata || true
	$(DOCKER_COMPOSE) up -d postgres
	@echo "$(GREEN)Database reset completed!$(NC)"

migration-up: ## Run database migrations
	@echo "$(YELLOW)Running migrations...$(NC)"
	$(DOCKER_COMPOSE) up flyway
	@echo "$(GREEN)Migrations completed!$(NC)"

migration-logs: ## Show migration logs
	$(DOCKER_COMPOSE) logs flyway

db-connect: ## Connect to PostgreSQL database
	$(DOCKER_COMPOSE) exec postgres psql -U myuser -d wallet

redis-cli: ## Connect to Redis CLI
	$(DOCKER_COMPOSE) exec redis redis-cli

# Development commands
dev: docker-up-build ## Start development environment
	@echo "$(GREEN)Development environment is ready!$(NC)"
	@echo "$(BLUE)Application: http://localhost:8080$(NC)"
	@echo "$(BLUE)PostgreSQL: localhost:5432$(NC)"
	@echo "$(BLUE)Redis: localhost:6379$(NC)"

dev-logs: ## Follow development logs
	$(DOCKER_COMPOSE) logs -f app postgres redis

stop: docker-down ## Stop development environment

# Health checks
health: ## Check service health
	@echo "$(YELLOW)Checking service health...$(NC)"
	@curl -f http://localhost:8080/health || echo "$(RED)Application health check failed$(NC)"
	@$(DOCKER_COMPOSE) exec postgres pg_isready -U myuser -d wallet || echo "$(RED)Database health check failed$(NC)"
	@$(DOCKER_COMPOSE) exec redis redis-cli ping || echo "$(RED)Redis health check failed$(NC)"

# Cleanup commands
clean: ## Clean build artifacts and Docker resources
	@echo "$(YELLOW)Cleaning up...$(NC)"
	rm -rf bin/
	rm -f coverage.out coverage.html
	go clean
	$(DOCKER_COMPOSE) down -v --remove-orphans
	docker system prune -f
	@echo "$(GREEN)Cleanup completed!$(NC)"

clean-docker: ## Clean all Docker resources
	@echo "$(YELLOW)Cleaning Docker resources...$(NC)"
	$(DOCKER_COMPOSE) down -v --remove-orphans
	docker system prune -a -f --volumes
	@echo "$(GREEN)Docker cleanup completed!$(NC)"

# Production commands
prod-build: ## Build for production
	@echo "$(YELLOW)Building for production...$(NC)"
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/$(APP_NAME) .
	@echo "$(GREEN)Production build completed!$(NC)"

# Git commands
git-hooks: ## Install git hooks
	@echo "$(YELLOW)Installing git hooks...$(NC)"
	cp scripts/pre-commit .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
	@echo "$(GREEN)Git hooks installed!$(NC)"

# Quick commands
quick-test: fmt vet test ## Quick development cycle: format, vet, test
	@echo "$(GREEN)Quick test cycle completed!$(NC)"

quick-deploy: clean docker-up-build health ## Quick deployment cycle
	@echo "$(GREEN)Quick deployment completed!$(NC)"

# Install tools
install-tools: ## Install development tools
	@echo "$(YELLOW)Installing development tools...$(NC)"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "$(GREEN)Development tools installed!$(NC)"