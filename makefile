
migration-up:
	migrate -path ./migrations/postgres -database 'postgres://user_name:password@0.0.0.0:5432/database?sslmode=disable' up
