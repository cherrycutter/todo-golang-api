build:
	docker-compose build

run:
	docker-compose up

test:
	go test -v ./internal/repos

swag:
	swag init -g ./cmd/app/main.go

migrate:
	migrate -path ./schema -database 'postgres://postgres:12345@localhost:5436/postgres?sslmode=disable' up





