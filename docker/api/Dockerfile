FROM samacs/coderoso:latest AS base

COPY . /cortito

WORKDIR /cortito/api

ENV GOOS=linux \
    CGO_ENABLED=0

RUN go mod download

FROM base AS builder

CMD dogo
