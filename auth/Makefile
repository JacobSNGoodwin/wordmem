.PHONY: migrate-create migrate-up migrate-down

N = 1

migrate-create:
	@echo "---Creating migration files---"
	migrate create -ext sql -dir migration -seq -digits 5 $(NAME)

migrate-up:
	migrate -source file://migration -database postgres://postgres:password@localhost/postgres?sslmode=disable up $(N)

migrate-down:
	migrate -source file://migration -database postgres://postgres:password@localhost/postgres?sslmode=disable down $(N)

migrate-force:
	migrate -source file://migration -database postgres://postgres:password@localhost/postgres?sslmode=disable force $(VERSION)