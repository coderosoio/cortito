FROM golang:latest AS base

COPY ./common /cortito/common
COPY ./account /cortito/account
COPY ./shortener /cortito/shortener
COPY ./api/go.mod /cortito/api/
COPY ./api/go.sum /cortito/api/

WORKDIR /cortito/api

RUN go mod download

FROM base AS builder

COPY . /cortito

WORKDIR /cortito/api

ENV GOOS=linux \
    CGO_ENABLED=0 \
    GO111MODULE=on

RUN go build -ldflags="-w -s" -o bin/api cmd/api/*.go

FROM alpine:latest

WORKDIR /cortito

COPY --from=builder /cortito/api/bin/api api

CMD ["./api"]
