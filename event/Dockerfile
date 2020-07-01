# Build
FROM golang:1.14

RUN mkdir -p /go/src \
    && mkdir -p /go/bin \
    && mkdir -p /go/pkg

WORKDIR $GOPATH/src

COPY go.sum .
COPY go.mod .
RUN go mod download
COPY . $GOPATH/src

#RUN go build

RUN apt-get update \
    && apt-get install -y mariadb-client 

EXPOSE 8081
CMD ["/bin/bash"]