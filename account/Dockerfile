FROM golang:1.14-alpine as builder

WORKDIR /account
COPY . .
RUN apk --no-cache add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
RUN ls
# RUN go build -o account .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o account .


FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /account
COPY --from=builder /account/account .
COPY --from=builder /account/rsa ./rsa
COPY --from=builder /account/migrations ./migrations
COPY --from=builder /account/account.env .
COPY --from=builder /account/start.sh .
COPY --from=builder /account/wait_for.sh .
COPY --from=builder /account/migrate.linux-amd64 ./migrate

RUN ls

EXPOSE 8080

CMD ["/account/account"]
ENTRYPOINT [ "/account/start.sh" ]