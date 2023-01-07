SHELL := /bin/bash
.DEFAULT_GOAL := default
.PHONY: all
BINARY_NAME=shortener

build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux-amd64 cmd/${BINARY_NAME}/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin-arm64 cmd/${BINARY_NAME}/main.go
	chmod +x bin/*

buildx:
	docker buildx build --platform="linux/amd64" -t loqutus/urlshortener:latest --push .
	docker buildx build --platform="linux/amd64" -t loqutus/urlshortener-test:latest --push -f ./Dockerfile-test .

get:
	go mod tidy

test:
	go test -count=1 -v -bench ./

install:
	cd deployments/urlshortener && helm dependency build
	helm install urlshortener ./deployments/urlshortener

uninstall:
	helm uninstall urlshortener || true

default: get build
