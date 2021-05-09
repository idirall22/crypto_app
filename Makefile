build:
	docker-compose up --build
up:
	docker-compose up -d
down:
	docker-compose down

docker-prune:
	docker rmi -f (docker images -f "dangling=true" -q)
