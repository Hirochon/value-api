FROM golang:1.19.1-alpine3.16

ARG MYSQL_USER
ARG MYSQL_PASSWORD
ARG MYSQL_DATABASE
ARG MYSQL_PORT
ARG MYSQL_HOST

ENV MYSQL_USER $MYSQL_USER
ENV MYSQL_PASSWORD $MYSQL_PASSWORD
ENV MYSQL_DATABASE $MYSQL_DATABASE
ENV MYSQL_PORT $MYSQL_PORT
ENV MYSQL_HOST $MYSQL_HOST

WORKDIR /go/src/value-api

COPY . /go/src/value-api/

RUN apk update && \
    apk add --no-cache git gcc musl-dev make

RUN go mod tidy && \
    go install github.com/google/wire/cmd/wire@v0.5.0 && \
    go build -o ./ ./...

EXPOSE 8600
