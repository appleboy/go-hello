# go-hello

hello world for go lang example.

[![Build Status](https://travis-ci.org/appleboy/go-hello.svg?branch=master)](https://travis-ci.org/appleboy/go-hello) [![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/go-hello)](https://goreportcard.com/report/github.com/appleboy/go-hello) [![codecov](https://codecov.io/gh/appleboy/go-hello/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/go-hello) [![Coverage Status](https://coveralls.io/repos/github/appleboy/go-hello/badge.svg?branch=master)](https://coveralls.io/github/appleboy/go-hello?branch=master) [![Build Status](https://drone.io/github.com/appleboy/go-hello/status.png)](https://drone.io/github.com/appleboy/go-hello/latest) [![Build Status](http://drone.wu-boy.com/api/badges/appleboy/go-hello/status.svg)](http://drone.wu-boy.com/appleboy/go-hello)

[![](http://badge-imagelayers.iron.io/appleboy/hello-world:latest.svg)](http://imagelayers.iron.io/?images=appleboy/hello-world:latest 'Get your own badge on imagelayers.iron.io')

## Getting Started

Install dependencies package.

```bash
$ go get -d
```

Start web server and default port is `8000`

```bash
$ go run hello-world.go
```

You can change default port via add `port` flag.

```bash
$ go run hello-world.go -port=8009
```

## Using Docker

Please install [Docker](https://docs.docker.com/machine/get-started/) first and run the following command to start server.

```bash
$ make all
```

Open your browser with `http://your_docker_ip:8000` and get the follwoing message.

```javascript
{
  "current_time":"2016-02-14T07:50:27.536032273Z",
  "text":"hello world"
}
```

Check docker process:

```bash
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                    NAMES
d044b5041e8e        hello-world         "/bin/sh -c /app/hell"   About an hour ago   Up About an hour    0.0.0.0:8000->8000/tcp   hello-production
```

Check docker images:

```bash
$ docker images
REPOSITORY                  TAG                 IMAGE ID            CREATED             VIRTUAL SIZE
hello-world                 latest              5cca15b3231f        43 seconds ago      15.98 MB
hello-world-build           latest              472296ce571d        47 seconds ago      282.5 MB
```

Rebuild your image.

```bash
$ make build
docker build -t "hello-world-build" -f Dockerfile.build .
Sending build context to Docker daemon 789.5 kB
Step 1 : FROM golang:1.6-alpine
 ---> 7ae766a8518d
Step 2 : MAINTAINER Bo-Yi Wu <appleboy.tw@gmail.com>
 ---> Using cache
 ---> 9cee8a14fec4
Step 3 : RUN apk --update add git
 ---> Using cache
 ---> d17daae20ba7
Step 4 : RUN mkdir -p /tmp/build
 ---> Using cache
 ---> cf511150ec0e
Step 5 : ADD hello-world.go /tmp/build/
 ---> Using cache
 ---> 02e933c9c453
Step 6 : WORKDIR /tmp/build
 ---> Using cache
 ---> a0d43ae462b5
Step 7 : RUN go get -d
 ---> Running in b02a8bd86c4f
 ---> 4946d82c8d0f
Removing intermediate container b02a8bd86c4f
Step 8 : RUN go build hello-world.go
 ---> Running in b88fcda825f1
 ---> 8e25074145ff
Removing intermediate container b88fcda825f1
Step 9 : CMD tar -czf - hello-world
 ---> Running in 2c6e38a708c7
 ---> 472296ce571d
Removing intermediate container 2c6e38a708c7
Successfully built 472296ce571d
docker run "hello-world-build" > build.tar.gz
docker build -t "hello-world" -f Dockerfile.dist .
Sending build context to Docker daemon 4.007 MB
Step 1 : FROM gliderlabs/alpine:3.3
 ---> 24e7cde1dbe9
Step 2 : MAINTAINER Bo-Yi Wu <appleboy.tw@gmail.com>
 ---> Using cache
 ---> 73e6d76da566
Step 3 : RUN mkdir /app
 ---> Using cache
 ---> c8c863a62423
Step 4 : ADD build.tar.gz /app/
 ---> 12ef80e2e0c6
Removing intermediate container 38ef062dc3bd
Step 5 : ENTRYPOINT /app/hello-world
 ---> Running in a4781a03cfd4
 ---> 07f53380e131
Removing intermediate container a4781a03cfd4
Step 6 : EXPOSE 8000
 ---> Running in 967f83e5c5f6
 ---> 5cca15b3231f
Removing intermediate container 967f83e5c5f6
Successfully built 5cca15b3231f
```

Restart your app server.

```bash
$ make server
docker rm -f hello-production
hello-production
docker run -d -p 8000:8000 --name "hello-production" "hello-world"
52723b9e29c2a422451e7a5e73fbd44be83c730006c02667140b43cf7f73c146
```

Deploy your image to [docker hub](https://hub.docker.com)

```bash
$ make deploy tag=4.0
docker tag "hello-world":latest "appleboy"/"hello-world":4.0
docker push "appleboy"/"hello-world":4.0
The push refers to a repository [docker.io/appleboy/hello-world] (len: 1)
80d1436c2e8f: Image already exists
c4bc05068eb6: Image already exists
c8c863a62423: Image already exists
24e7cde1dbe9: Image already exists
4.0: digest: sha256:b5b6bd28e8a4ced89d4e38fe05b9fa4d8f3483f23f23ef1ab0f66b23ffdbcf0e size: 9507
```

## Testing

```bash
$ make test
go get -t -v ./...
go test -v
PASS
ok    github.com/appleboy/go-hello  0.012s
```
