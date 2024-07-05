.PHONY: deps
deps:
	@echo "==> Setup: Tidying modules"
	go mod tidy
	@echo "==> Setup: Getting dependencies"
	go mod download

.PHONY: generate
generate:
	@echo "==> Go Generating"
	@go generate ./...
	@echo "Running swag init"
	@swag init -d app/gateway/http -g router.go -o docs/swagger

.PHONY: start
start:
	@echo "==> Starting server"
	@go run cmd/api/main.go

.PHONY: build
build:
	@echo "==> Building server"
	@go build -o bin/api cmd/api/main.go