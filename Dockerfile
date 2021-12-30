#stage 1
FROM golang:1.17.0-alpine3.14 AS builder

RUN mkdir /app
ADD . /app

WORKDIR /app

RUN go clean --modcache
RUN go build -o app/main

#stage 2
FROM alpine:3.14
WORKDIR /root/

COPY --from=builder /app/config.json .
COPY --from=builder /app/app/main .

EXPOSE 9090

CMD ["./main"]