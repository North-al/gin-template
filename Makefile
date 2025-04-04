version ?= latest

run:
	go run cmd/main.go

build:
	docker build -t gin-template:$(version) .

save:
	docker save -o ./docker-build/restaurant-$(version).tar restaurant:$(version)

dev:
	air -c .air.toml

test:
	go test ./...

swag:
	swag init -g ./cmd/main.go

gen:
	go run cmd/gen/main.go