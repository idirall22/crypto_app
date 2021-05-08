FROM golang:1.14 as builder

WORKDIR /account
COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o account cmd/server/main.go
RUN go build -o account cmd/server/main.go

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /account/account .

EXPOSE 8081

CMD ["./account"] 