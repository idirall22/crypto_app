include account.env

DOCKER_POSTGRES_NAME=crypto

rsa-genrate:
	openssl genrsa -out rsa/key.pem 2048 && openssl rsa -in rsa/key.pem -pubout -out rsa/public.pem

# mock services
mocks:
	mockery --dir adapters/repository/ --name IRepository --output adapters/repository/postgres/mock --exported
	mockery --dir adapters/email/ --name IEmail --output adapters/email/mock --exported
	mockery --dir service/ --name IService --output service/mock --exported

# start a postgresql docker container used to run adapters tests.
up:
	docker run -d -p ${DB_PORT}:${DB_PORT} -e POSTGRES_PASSWORD=${DB_PASSWORD} --name ${DOCKER_POSTGRES_NAME} postgres:13-alpine

# delete the postgresql docker container used to run adapters tests.
down:
	docker rm -f ${DOCKER_POSTGRES_NAME}

# create database
createdb:
	docker exec -it ${DOCKER_POSTGRES_NAME} sh -c "createdb -U postgres $(DB_NAME)"

# drop database
dropdb:
	docker exec -it ${DOCKER_POSTGRES_NAME} sh -c "dropdb -U postgres --if-exists $(DB_NAME)"

# connect to database
connect:
	docker exec -it ${DOCKER_POSTGRES_NAME} sh -c "psql -U postgres -d $(DB_NAME) copy workflow_facilitator_types FROM '/workflow_facilitator_types.csv' DELIMITER ',' CSV"

# migrate the sql code to the database
migrate: dropdb createdb
	migrate -path migrations/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up

test_adapters: dropdb createdb migrate
	go test -v -cover -count=1 ./adapters/repository/postgres/test/...

test_service:
	go test -v -cover -count=1 ./service/test/...

test_port:
	go test -v -cover -count=1 ./port/test/...

run:
	go run cmd/server/main.go

build:
	docker build -t account:latest .
	docker rmi -f $(docker images -f "dangling=true" -q)
	
.PHONY: up down createdb dropdb connect migrate sqlc test_services test_adapters

