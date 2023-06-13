-include .env
export

start:
	go run main.go

seed:
	go run main.go seed

postgres:
	@PGPASSWORD=${DB_PASS} psql -h ${DB_HOST} -U ${DB_USER} -p ${DB_PORT} ${DB_NAME}

migrate:
	@goose -dir migrations postgres "user=${DB_USER} port=${DB_PORT} password=${DB_PASS} dbname=${DB_NAME} host=${DB_HOST} sslmode=disable" up

status:
	@goose -dir migrations postgres "user=${DB_USER} port=${DB_PORT} password=${DB_PASS} dbname=${DB_NAME} host=${DB_HOST} sslmode=disable" status
