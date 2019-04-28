FROM golang:latest AS builder

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
