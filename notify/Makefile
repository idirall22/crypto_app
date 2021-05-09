include .env
include notify.env

rabbit-up:
	docker run \
	--hostname crypto-rabbit \
	-e RABBITMQ_DEFAULT_USER=${RABBITMQ_USER} \
	-e RABBITMQ_DEFAULT_PASS=${RABBITMQ_PASSWORD} \
	-p 5672:5672 -p 15672:15672 \
	--name crypto-rabbit rabbitmq:3

rabbit-down:
	docker rm -f crypto-rabbit

server:
	go run main.go

# go run play/main.go
playr:
	go run cmd/server/server.go