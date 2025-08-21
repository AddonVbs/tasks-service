MIGRATE=migrate -path ./migrations -database "postgres://admin:2311@localhost:5432/postgres?sslmode=disable"

.PHONY: migrate migrate-down

DB_URL=postgres://admin:2311@localhost:5433/postgres?sslmode=disable

migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down 1

migrate-force:
	@if [ -z "$(version)" ]; then echo "Usage: make migrate-force version=1234567890"; exit 1; fi
	migrate -path ./migrations -database "$(DB_URL)" force $(version)

migrate-drop:
	migrate -path ./migrations -database "$(DB_URL)" drop -f
