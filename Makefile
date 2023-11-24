include .env

migrate-up:
	migrate -database ${DB_SOURCE} -path db/migrations up

migrate-down:
	migrate -database ${DB_SOURCE} -path db/migrations down --all

populate:
	psql ${DB_SOURCE} -f internal/db/migrations/0000001_init_schema.up.sql

down:
	docker-compose down --volumes && docker volume prune -f

up:
	docker-compose up -d

sqlc:
	rm -rf internal/repository
	sqlc generate

start:
	make up
	docker exec orders-psql sh /tmp/health.sh
	go run cmd/http_viewer/http_viewer.go

restart:
	make down
	make up
	docker exec orders-psql sh /tmp/health.sh
	make migrate-up
	make start

.PHONY: migrate-up migrate-down populate down up sqlc run
