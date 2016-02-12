#!/bin/sh

output() {
  printf "\E[0;33;40m"
  echo $1
  printf "\E[0m"
}

output "Remove hello container."
docker rm -f hello-production
output "Start hello container."
docker run -d -p 8000:8000 --name hello-production hello-world
