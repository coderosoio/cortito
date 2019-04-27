FROM golang:latest

COPY go.mod go.sum /

RUN go get -u github.com/samacs/dogo

RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN go mod download
