run:
	go run cmd/main.go

build:
	go build -o bin/app cmd/main.go

migrate:
	go run internal/data/migrations/migrate.go

test:
	go test ./...
