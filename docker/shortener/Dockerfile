FROM samacs/coderoso:latest

COPY . /cortito

WORKDIR /cortito/shortener

ENV GOOS=linux \
    CGO_ENABLED=0

CMD dogo
