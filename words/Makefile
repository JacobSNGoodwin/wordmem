.PHONY: migrate-create migrate-up migrate-down migrate-force

PWD = $(shell pwd)
MPATH = $(PWD)/src/migrations

# docker command to run migrations
# MIGRATE = docker run --rm --name migrateTmp -v $(MPATH):/migrations --network host migrate/migrate
N = 1

migrate-create:
	@echo "---Creating migration files---"
	migrate create -ext sql -dir $(MPATH) -seq -digits 5 $(NAME)
	# $(MIGRATE) create -ext sql -dir migrations -seq -digits 5 $(NAME)
	# docker cp migrateTmp:migrations $(MPATH)

migrate-up:
	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:5433/postgres?sslmode=disable up $(N)

migrate-down:
	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:5433/postgres?sslmode=disable down $(N)

migrate-force:
	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:5433/postgres?sslmode=disable force $(VERSION)