include .env

build-all:
	go build -o bin/social ./cmd/social
stop-all:
	docker compose down

run-all:  build-all
	$(MAKE) stop-all
	docker compose down
	docker compose build
	docker compose up -d social-db-0
	$(MAKE) migrate-up
	docker compose up -d --force-recreate --build social

migrate-down:
	sleep 2

	GOOSE_MIGRATION_DIR=${GOOSE_MIGRATION_DIR} \
	GOOSE_DRIVER=${GOOSE_DRIVER} \
	GOOSE_DBSTRING=${GOOSE_DBSTRING} \
	goose reset

migrate-up:
	sleep 2

	GOOSE_MIGRATION_DIR=${GOOSE_MIGRATION_DIR} \
	GOOSE_DRIVER=${GOOSE_DRIVER} \
	GOOSE_DBSTRING=${GOOSE_DBSTRING} \
	goose up

tidy:
	go mod tidy