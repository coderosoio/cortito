FROM golang:latest AS builder

COPY . /cortito

WORKDIR /cortito/shortener

ENV GOOS=linux \
    CGO_ENABLED=0 \
    GO111MODULE=on \
    APP_ENV=production

RUN go build -ldflags="-w -s" -o bin/shortener cmd/shortener/*.go

FROM alpine:latest

WORKDIR /cortito

COPY --from=builder /cortito/shortener/bin/shortener shortener

ENV APP_ENV=production

CMD ["./shortener"]
