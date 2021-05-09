rabbit-up:
	docker run \
	--hostname crypto-rabbit \
	-e RABBITMQ_DEFAULT_USER=${RABBITMQ_USER} \
	-e RABBITMQ_DEFAULT_PASS=${RABBITMQ_PASSWORD} \
	-p 5672:5672 -p 15672:15672 \
	--name crypto-rabbit rabbitmq:3

rabbit-down:
	docker rm -f crypto-rabbit

build:
	docker-compose up --build
up:
	docker-compose up -d
down:
	docker-compose down

prune:
	docker rmi -f $(docker images -f "dangling=true" -q)

docker-account:
	docker build -f account/Dockerfile .

docker-notify:
	docker build -f notify/Dockerfile .