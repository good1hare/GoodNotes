migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'migrate_name'
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -path migrations -database 'postgres://postgres:password@0.0.0.0:5433/postgres?sslmode=disable' up
.PHONY: migrate-up