
dev:
	go run main.go
build:
	go build -o ./dist/app
migrate:
	go run ./development/migrate/migrate.go
seed:
	go run ./development/seed/seed.go