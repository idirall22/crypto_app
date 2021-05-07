include example.env

DOCKER_POSTGRES_NAME=crypto

rsa-genrate:
	openssl genrsa -out key.pem 2048 && openssl rsa -in key.pem -pubout -out public.pem

# mock services
mocks:
	mockery --dir account/adapters/repository/ --name IRepository --output account/adapters/repository/postgres/mock --exported
	mockery --dir account/adapters/email/ --name IEmail --output account/adapters/email/mock --exported
	mockery --dir account/service/ --name IService --output account/service/mock --exported

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
	migrate -path account/migrations/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up

test_adapters: dropdb createdb migrate
	go test -v -cover -count=1 ./account/adapters/repository/postgres/test/...

test_service:
	go test -v -cover -count=1 ./account/service/test/...

test_port: dropdb createdb migrate
	go test -v -cover -count=1 ./ports/http/test/...

.PHONY: up down createdb dropdb connect migrate sqlc test_services test_adapters

