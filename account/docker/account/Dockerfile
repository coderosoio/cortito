FROM golang:latest AS base

COPY ./common /cortito/common
COPY ./account/go.mod /cortito/account/
COPY ./account/go.sum /cortito/account/

WORKDIR /cortito/account

RUN go mod download

FROM base AS builder

COPY . /cortito

WORKDIR /cortito/account

ENV GOOS=linux \
    CGO_ENABLED=0 \
    GO111MODULE=on

RUN go build -ldflags="-w -s" -o bin/account cmd/account/*.go

FROM alpine:latest

WORKDIR /cortito

COPY --from=builder /cortito/account/bin/account account

CMD ["./account"]
