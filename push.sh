#!/usr/bin/env bash

if [ $# -ne 2 ]; then
  >&2 echo "Usage: $0 <author> <tag>"
  exit 1
fi

docker tag hello-world:latest $1/hello-world:$2
docker push $1/hello-world:$2
