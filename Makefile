
VERSION         :=      $(shell cat ./VERSION)
IMAGE_NAME      :=      fallenstedt/test-grpc

all: install

install:
	echo "Installing go modules..." && \
	go mod download && \
	echo "Completed" 

build:
	go build main.go

test:
	go test ./src/... -v

image:
	docker build -t fallenstedt/gin-example .

release:
	git tag -a $(VERSION) -m "Release" || true
	git push origin $(VERSION)

.PHONY: install test fmt build release
