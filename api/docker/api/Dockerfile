FROM golang:latest AS builder

COPY . /cortito

WORKDIR /cortito/api

ENV GOOS=linux \
    CGO_ENABLED=0 \
    GO111MODULE=on

RUN go build -ldflags="-w -s" -mod=vendor -o bin/api cmd/api/*.go

FROM alpine:latest

WORKDIR /cortito

COPY --from=builder /cortito/api/bin/api api

CMD ["./api"]
