FROM golang:1.14-alpine as builder

WORKDIR /notify
COPY wait_for.sh .
COPY rsa ./rsa
COPY ./notify .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o notify .


FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /notify
COPY --from=builder /notify/notify .
COPY --from=builder /notify/rsa ./rsa
COPY --from=builder /notify/notify.env . 
COPY --from=builder /notify/start.sh .
COPY --from=builder /notify/wait_for.sh .

EXPOSE 8081

CMD ["/notify/notify"]
ENTRYPOINT [ "/notify/start.sh" ]