.PHONY: all

DEPS := $(wildcard *.go)
BUILD_IMAGE := "hello-world-build"
PRODUCTION_IMAGE := "hello-world"
PRODUCTION_NAME := "hello-production"
DEPLOY_ACCOUNT := "appleboy"

all: build server

build:
	docker build -t $(BUILD_IMAGE) -f Dockerfile.build .
	docker run $(BUILD_IMAGE) > build.tar.gz
	docker build -t $(PRODUCTION_IMAGE) -f Dockerfile.dist .

server:
	-docker rm -f hello-production
	-docker run -d -p 8000:8000 --name $(PRODUCTION_NAME) $(PRODUCTION_IMAGE)

deploy:
ifeq ($(tag),)
	@echo "Usage: make $@ tag=<tag>"
	@exit 1
endif
	docker tag $(PRODUCTION_IMAGE):latest $(DEPLOY_ACCOUNT)/$(PRODUCTION_IMAGE):$(tag)
	docker push $(DEPLOY_ACCOUNT)/$(PRODUCTION_IMAGE):$(tag)

hello: ${DEPS}
	GO15VENDOREXPERIMENT=1 go build

test:
	go get -d
	go test

clean:
	rm -rf build.tar.gz
	go clean

dist-clean: clean
	-docker rmi -f $(BUILD_IMAGE)
	-docker rm -f $(PRODUCTION_NAME)
	-docker rmi -f $(PRODUCTION_IMAGE)
