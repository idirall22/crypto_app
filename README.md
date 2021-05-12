## Cryptoapp

A simple applicaton that allows users to have wallets and send money to other users.
The architecture used to develop this app is the **Hexagonal Architecture**.
There are 2 main module the **account** and the **notify** module.

## Database design

## Run the application
### Local
0. Run Account tests:
- 

To run the application localy, you have to use docker-compose.yaml file
1. run `./init.sh` this allows to add the http://cryptoapp.com url to the `/etc/hosts` file
2. to be able to send email you hace to use a gmail account and set the env variables
`GMAIL_PASSWORD` and `GMAIL_EMAIL`
example:
- export GMAIL_PASSWORD=secret
- export GMAIL_EMAIL=email@gmail.com
`note: if the gmail credentials are not set the application can still work`
3. use postman to test the application:
    - import `cryptoapp_collection.postman_collection.json` in postman
