
GOOGLE_CLOUD_PROJECT_ID := gin-micro
VERSION         		:= $(shell cat ./VERSION)

PROJECT_PREFIX			:= fallenstedt
PROJECT_NAME			:= gin-mongodb-starter

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
	docker build -t $(PROJECT_PREFIX)/$(PROJECT_NAME):v$(VERSION) .

gcr_tag:
	docker tag $(PROJECT_PREFIX)/$(PROJECT_NAME):v$(VERSION) gcr.io/$(GOOGLE_CLOUD_PROJECT_ID)/$(PROJECT_NAME):v$(VERSION)

gcr_release:
	docker push  gcr.io/$(GOOGLE_CLOUD_PROJECT_ID)/$(PROJECT_NAME):v$(VERSION)

release:
	git tag -a $(VERSION) -m "Release" || true
	git push origin $(VERSION)

.PHONY: install test fmt build release
