
include .env
build:
	go build -o bin/social ./cmd/social
run:
	@go run ./cmd/social ./cmd/social

migrate-up:
	migrate -path ./migrations -database 'mysql://root:password@tcp(localhost:${DB_EXTERNAL_PORT})/social?query' up
migrate-down:
	migrate -path ./migrations -database 'mysql://root:password@tcp(localhost:${DB_EXTERNAL_PORT})/social?query' down
migrate-refresh:
	make migrate-down && make migrate-up