rsa-genrate:
	openssl genrsa -out rsa/key.pem 2048 && openssl rsa -in rsa/key.pem -pubout -out rsa/public.pem

rabbit-up:
	docker run \
	--hostname crypto-rabbit \
	-e RABBITMQ_DEFAULT_USER=user \
	-e RABBITMQ_DEFAULT_PASS=password \
	-p 5672:5672 -p 15672:15672 \
	--name crypto-rabbit rabbitmq:3

rabbit-down:
	docker rm -f crypto-rabbit

build:
	docker-compose build
up:
	echo ${GMAIL_PASSWORD}
	GMAIL_PASSWORD=${GMAIL_PASSWORD} docker-compose up
down:
	docker-compose down

prune:
	docker rmi -f $(docker images -f "dangling=true" -q)

docker-account:
	docker build -f account/Dockerfile .

docker-notify:
	docker build -f notify/Dockerfile .