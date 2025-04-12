.PHONY: build run test migrate-up migrate-down migrate-create

build:
	go build -o bin/smtest ./cmd

run:
	go run ./cmd

test:
	go test ./...

migrate-up:
	migrate -path migrations/postgres -database "postgres://postgres:postgres@localhost:5432/smtest?sslmode=disable" up

migrate-down:
	migrate -path migrations/postgres -database "postgres://postgres:postgres@localhost:5432/smtest?sslmode=disable" down

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations/postgres -seq $$name