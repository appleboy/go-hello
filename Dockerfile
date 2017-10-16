FROM golang:1.9-alpine3.6

MAINTAINER Bo-Yi Wu <appleboy.tw@gmail.com>

ADD . /go/src/github.com/appleboy/go-hello

RUN go install github.com/appleboy/go-hello

EXPOSE 8000
CMD ["/go/bin/go-hello"]
