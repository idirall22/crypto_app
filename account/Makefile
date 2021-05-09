include account.env

DOCKER_POSTGRES_NAME=cryptoapp


# mock services
mocks:
	mockery --dir adapters/repository/ --name IRepository --output adapters/repository/postgres/mock --exported
	mockery --dir adapters/event --name IEventStore --output adapters/event/amqp/mock --exported
	mockery --dir service/ --name IService --output service/mock --exported

# start a postgresql container
up:
	docker run -d -p ${DB_PORT}:${DB_PORT} -e POSTGRES_PASSWORD=${DB_PASSWORD} --name ${DOCKER_POSTGRES_NAME} postgres:13-alpine

# delete the postgresql docker
down:
	docker rm -f ${DOCKER_POSTGRES_NAME}

# create database
createdb:
	docker exec -it ${DOCKER_POSTGRES_NAME} sh -c "createdb -U postgres $(DB_NAME)"

# drop database
dropdb:
	docker exec -it ${DOCKER_POSTGRES_NAME} sh -c "dropdb -U postgres --if-exists $(DB_NAME)" || true

# start rabbitmq container
rabbit-up:
	docker run \
	--hostname crypto-rabbit \
	-e RABBITMQ_DEFAULT_USER=${RABBITMQ_USER} \
	-e RABBITMQ_DEFAULT_PASS=${RABBITMQ_PASSWORD} \
	-p 5672:5672 -p 15672:15672 \
	--name crypto-rabbit rabbitmq:3

# delete rabbitmq container
rabbit-down:
	docker rm -f crypto-rabbit

# connect to database
connect:
	docker exec -it ${DOCKER_POSTGRES_NAME} sh -c "psql -U postgres -d $(DB_NAME) copy workflow_facilitator_types FROM '/workflow_facilitator_types.csv' DELIMITER ',' CSV"

# migrate the sql code to the database
migrate:
	migrate -path migrations/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up

test_adapters: dropdb createdb migrate
	go test -v -cover -count=1 ./adapters/repository/postgres/test/...

test_service:
	go test -v -cover -count=1 ./service/test/...

test_port:
	go test -v -cover -count=1 ./port/test/...

run:
	go run cmd/server/main.go
	
.PHONY: up down createdb dropdb connect migrate sqlc test_services test_adapters

