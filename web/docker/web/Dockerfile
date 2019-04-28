FROM golang:latest AS builder

COPY . /cortito

WORKDIR /cortito/web

ENV GOOS=linux \
    CGO_ENABLED=0 \
    GO111MODULE=on

RUN go build -ldflags="-w -s" -o bin/web cmd/web/*.go

FROM alpine:latest

WORKDIR /cortito

COPY --from=builder /cortito/web/bin/web web

CMD ["./web"]
