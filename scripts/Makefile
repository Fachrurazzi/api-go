setup:
	@echo "Setup Postgres..\n"
	@echo "\nCREATING database"
	psql -c "DROP DATABASE IF EXISTS logic"
	psql -c "DROP DATABASE IF EXISTS logic_test"
	psql -c "DROP USER IF EXISTS testing"
	psql -c "CREATE DATABASE logic;"
	psql -c "CREATE DATABASE logic_test;"
	@echo "\nCREATING user with password"
	psql -c "CREATE USER testing WITH PASSWORD 'test' VALID UNTIL 'December 20 2020';"
	psql -c "ALTER USER testing WITH SUPERUSER"
	@echo "\nGRANT ALL database to user"
	psql -c "GRANT ALL PRIVILEGES ON DATABASE logic to testing;"
	psql -c "GRANT ALL PRIVILEGES ON DATABASE logic_test to testing;"
	@echo "\nFinish Setup Postgres :)"
db_seed:
	@echo "Seeding....\n"
	psql -U testing -h localhost -d logic -c \
	"CREATE OR REPLACE FUNCTION trigger_set_timestamp() RETURNS TRIGGER AS 'BEGIN NEW.updated_at = NOW(); RETURN NEW; END;' LANGUAGE plpgsql;"
	export DB_DEV_USERNAME=testing DB_DEV_PASSWORD=test DB_DEV_NAME=logic && \
	sql-migrate up -config=../configs/seedconfig.yml -env="development"
db_seed_test:
	@echo "Seeding....\n"
	psql -U testing -h localhost -d logic_test -c \
	"CREATE OR REPLACE FUNCTION trigger_set_timestamp() RETURNS TRIGGER AS 'BEGIN NEW.updated_at = NOW(); RETURN NEW; END;' LANGUAGE plpgsql;"
	export DB_TEST_USERNAME=testing DB_TEST_PASSWORD=test DB_TEST_NAME=logic_test && \
	sql-migrate up -config=../configs/seedconfig.yml -env="test"
db_seed_undo:
	@echo "Seeding....\n"
	export DB_DEV_USERNAME=testing DB_DEV_PASSWORD=test DB_DEV_NAME=logic && \
	sql-migrate down -config=../configs/seedconfig.yml -env="development"
db_migrate:
	@echo "Migration....\n"
	psql -U testing -h localhost -d logic -c \
	"CREATE OR REPLACE FUNCTION trigger_set_timestamp() RETURNS TRIGGER AS 'BEGIN NEW.updated_at = NOW(); RETURN NEW; END;' LANGUAGE plpgsql;"
	export DB_DEV_USERNAME=testing DB_DEV_PASSWORD=test DB_DEV_NAME=logic && \
	sql-migrate up -config=../configs/dbconfig.yml -env="development"
db_migrate_test:
	@echo "Migration....\n"
	psql -U testing -h localhost -d logic_test -c \
	"CREATE OR REPLACE FUNCTION trigger_set_timestamp() RETURNS TRIGGER AS 'BEGIN NEW.updated_at = NOW(); RETURN NEW; END;' LANGUAGE plpgsql;"
	export DB_TEST_USERNAME=testing DB_TEST_PASSWORD=test DB_TEST_NAME=logic_test && \
	sql-migrate up -config=../configs/dbconfig.yml -env="test"
db_undo:
	@echo "Migration Undo....\n"
	export DB_DEV_USERNAME=testing DB_DEV_PASSWORD=test DB_DEV_NAME=logic && \
	sql-migrate down -config=../configs/dbconfig.yml -env="development"
db_test_dev:
	@echo "Testing Migrate before applied\n"
	export DB_DEV_USERNAME=testing DB_DEV_PASSWORD=test DB_DEV_NAME=logic && \
	sql-migrate up -config=../configs/dbconfig.yml -env="development" -dryrun
db_status:
	export DB_DEV_USERNAME=testing DB_DEV_PASSWORD=test DB_DEV_NAME=logic && \
	sql-migrate status -config=../configs/dbconfig.yml -env="development"
dev:
	go run ../main.go
prod:
	../api
test:
	psql -U testing -h localhost -d logic_test -c "DELETE from users_quotes"
	psql -U testing -h localhost  -d logic_test -c "DELETE from users"
	go test -count=1 -v ../test
	psql -U testing -h localhost -d logic_test -c "SELECT * from users_quotes"
	psql -U testing -h localhost -d logic_test -c "SELECT * from users"