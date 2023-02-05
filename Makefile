SHELL := /bin/bash
.DEFAULT_GOAL := default
.PHONY: all
BINARY_NAME=shortener
IMAGE_TAG=$(shell git describe --tags --always)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
ORG_PREFIX := loqutus

build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux-amd64 cmd/${BINARY_NAME}/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin-arm64 cmd/${BINARY_NAME}/main.go
	chmod +x bin/*

buildx:
	docker build -t loqutus/urlshortener:$(IMAGE_TAG)-$(GIT_COMMIT)  .
	docker push loqutus/urlshortener:$(IMAGE_TAG)-$(GIT_COMMIT)
	docker build -t loqutus/urlshortener-test:$(IMAGE_TAG)-$(GIT_COMMIT) -f ./Dockerfile-test .
	docker tag loqutus/urlshortener-test:$(IMAGE_TAG)-$(GIT_COMMIT) loqutus/urlshortener-test:latest
	docker push loqutus/urlshortener-test:$(IMAGE_TAG)-$(GIT_COMMIT)
	docker push loqutus/urlshortener-test:latest

get:
	go mod tidy

test:
	go test -v ./...

kubetest:
	kubectl apply -n test -f ./deployments/test/job.yaml
	kubectl wait --for=condition=complete --timeout=60s -n test job/urlshortener-test
	status=$(kubectl get jobs urlshortener-test -n test -o jsonpath='{.status.conditions[?(@.type=="Complete")].status}'); if [ "$status" == "True" ]; then exit 0; else exit 1; fi	

upgrade-prod:
	helm repo add bitnami https://charts.bitnami.com/bitnami
	cd deployments/urlshortener && helm dependency build
	helm upgrade -n prod --wait --timeout 2m --values ./deployments/urlshortener/values.yaml --set image.tag=$(GIT_TAG)-$(GIT_COMMIT) urlshortener ./deployments/urlshortener

uninstall-test:
	kubectl delete -n test -f ./deployments/test/job.yaml || true
	kubectl delete -n test -f ./deployments/test/postgres.yaml || true

clean:
	docker system prune -a -f

logs:
	kubectl logs -n test job/urlshortener-test

default: get build
