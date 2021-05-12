## Cryptoapp

A simple app that allows users to have wallets and send money to other users.
The application contains 2 main services: the `account` and `notify` service, each was developed using **Hexagonal Architecture**. The services communicate with each other using **Event driven Architecture** in this case, the account service communicates with the notification via Rabbitmq.
And in front of the services, there is an `Nginx` and its role is to be an API gateway, so we can make a request to a single service (Nginx) then it will automatically redirect the request to a service.

![architecture](https://raw.githubusercontent.com/idirall22/crypto_app/main/cryptoapp.png)

## Database design
![database](https://raw.githubusercontent.com/idirall22/crypto_app/main/crypto.png)

## Run the application
I recorded a video that you can follow to run the application localy [link](https://youtu.be/8NNINTOq8GE)

To run the application localy, you have to use docker-compose.yaml file.
1. Run `./init.sh` this allows to add the http://cryptoapp.com url to the `/etc/hosts` file.
2. To be able to send emails you have to use a Gmail account and set the env variables
`GMAIL_PASSWORD` and `GMAIL_EMAIL`
```js
- example:
    - export `GMAIL_PASSWORD=password`
    - export `GMAIL_EMAIL=email@gmail.com`
note: if the gmail credentials are not set the application will not sent emails
```
3. Build the docker-compose file using `make build`
4. Run the docker-compose file using `make up`
5. import `cryptoapp_collection.postman_collection.json` in postman
6. Inside postman add environment group
7. Run requests:
    * register user1
    * register user2
    * activate user1 account
    * activate user2 account
    * login user1
    * login user2
    * copy the `access_token` user1 and replace the one inside the `curl_websocket.txt` file then copy the first block and run it inside the terminal, this allows to connect to websocket.
    * copy the `access_token` user2 and replace the one inside the `curl_websocket.txt` file then copy the first block and run it inside the terminal, this allows to connect to websocket.
    * list user1 wallets
    * list user2 wallets
    * send money user1
    * list transactions
8. Run `make down` to stop the application.
