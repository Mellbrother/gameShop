FROM golang:latest

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . ./src

RUN apt-get update \
    && apt-get install -y mariadb-client

EXPOSE 8080