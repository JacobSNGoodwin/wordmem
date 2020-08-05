.PHONY: migrate-create migrate-up migrate-down


migrate-create:
	@echo "---Creating migration files---"
	migrate create -ext sql -dir migration -seq -digits 5 $(NAME)

migrate-up:
	migrate -source file://migration -database postgres://postgres:password@localhost/postgres?sslmode=disable up

migrate-down:
	migrate -source file://migration -database postgres://postgres:password@localhost/postgres?sslmode=disable down

migrate-force:
	migrate -source file://migration -database postgres://postgres:password@localhost/postgres?sslmode=disable force $(VERSION)