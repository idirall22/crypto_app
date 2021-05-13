GCP_PROJECT_ID=gateway-282214

rsa-genrate:
	openssl genrsa -out rsa/key.pem 2048 && openssl rsa -in rsa/key.pem -pubout -out rsa/public.pem

cert:
# 1. Generate CA's private key and self-signed certificate
	openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout nginx/certs/ca-key.pem -out nginx/certs/ca-cert.pem -subj "/C=DZ/ST=Oran/L=Oran/O=Crypto app/OU=Crypto/CN=*.crypto.com/emailAddress=cryptoapp@email.com"

# echo "CA's self-signed certificate"
	openssl x509 -in nginx/certs/ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
	openssl req -newkey rsa:4096 -nodes -keyout nginx/certs/server-key.pem -out nginx/certs/server-req.pem -subj "/C=DZ/ST=Oran/L=Oran/O=Crypto app/OU=Crypto/CN=*.crypto.com/emailAddress=cryptoapp@email.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
	openssl x509 -req -in nginx/certs/server-req.pem -days 60 -CA nginx/certs/ca-cert.pem -CAkey nginx/certs/ca-key.pem -CAcreateserial -out nginx/certs/server-cert.pem -extfile nginx/certs/server-ext.cnf

# echo "Server's signed certificate"
	openssl x509 -in nginx/certs/server-cert.pem -noout -text

build:
	docker-compose build

up:
	GMAIL_EMAIL=${GMAIL_EMAIL} GMAIL_PASSWORD=${GMAIL_PASSWORD} docker-compose up

down:
	docker-compose down

prune:
	docker rmi -f $(docker images -f "dangling=true" -q)

docker-account	:
	docker build -f account/Dockerfile -t gcr.io/${GCP_PROJECT_ID}/account .

docker-notify:
	docker build -f notify/Dockerfile -t gcr.io/${GCP_PROJECT_ID}/notify .

docker-login:
	cat sa.json | docker login -u _json_key --password-stdin https://gcr.io

gcp-push:
	docker push "gcr.io/${GCP_PROJECT_ID}/notify"
	docker push "gcr.io/${GCP_PROJECT_ID}/account"

gcp-pull:
	docker pull "gcr.io/${GCP_PROJECT_ID}/notify"
	docker pull "gcr.io/${GCP_PROJECT_ID}/account"

secrets:
	kubectl create secret generic gmail --from-literal=GMAIL_PASSWORD=${GMAIL_PASSWORD} --from-literal=GMAIL_EMAIL=${GMAIL_EMAIL}
	kubectl create secret generic password --from-literal=RABBITMQ_PASSWORD=${RABBITMQ_PASSWORD} --from-literal=DB_PASSWORD=${DB_PASSWORD}

tiller:
	kubectl create serviceaccount --namespace kube-system tiller
	kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
	helm init --service-account tiller --upgrade