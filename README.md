# go-hello

hello world for go lang example.

[![Build Status](https://travis-ci.org/appleboy/go-hello.svg?branch=master)](https://travis-ci.org/appleboy/go-hello) [![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/go-hello)](https://goreportcard.com/report/github.com/appleboy/go-hello)

## Getting Started

Please install [Docker](https://docs.docker.com/machine/get-started/) first and run the following command to start server.

```bash
$ make all
```

Open your browser with `http://your_docker_ip:8000` and get the follwoing message.

```javascript
{
  "currentTime":"2016-02-14T07:50:27.536032273Z",
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
hello-world                 latest              80d1436c2e8f        About an hour ago   15.98 MB
appleboy/hello-world        latest              80d1436c2e8f        About an hour ago   15.98 MB
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
make test
```
