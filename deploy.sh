#!/bin/sh

output() {
  printf "\E[0;33;40m"
  echo $1
  printf "\E[0m"
}

output "Stop container."
docker stop hello-production
output "Remove container."
docker rm -f hello-production
output "Start container."
docker run -d -p 8000:8000 --name hello-production hello-world
