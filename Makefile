.DEFAULT_GOAL := build
.PHONY:fmt vet build
fmt:
	@go fmt ./...
vet: fmt
	@go vet ./...
build: vet
	@go build -o bin/main
run: build
	./bin/main

kill-port:
	@sudo kill -9 `sudo lsof -t -i:3000`
dev:
	@modd
db:
	@docker compose up -d
migrate:
	@docker compose up migrate
