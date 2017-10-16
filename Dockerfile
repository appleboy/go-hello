FROM golang:1.9.1-alpine3.6

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>"

ADD . /go/src/github.com/appleboy/go-hello

RUN go install github.com/appleboy/go-hello

EXPOSE 8000
CMD ["/go/bin/go-hello"]
