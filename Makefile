SHELL := /bin/bash
.DEFAULT_GOAL := default
.PHONY: all
BINARY_NAME=shortener

build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux-amd64 cmd/${BINARY_NAME}/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin-arm64 cmd/${BINARY_NAME}/main.go
	chmod +x bin/*

buildx:
	docker build -t loqutus/urlshortener:latest .
	docker push loqutus/urlshortener:latest
	docker build -t loqutus/urlshortener-test:latest -f ./Dockerfile-test .
	docker push loqutus/urlshortener-test:latest

get:
	go mod tidy

test:
	go test -v ./...

kubetest:
	kubectl apply -n test -f ./deployments/test/pod.yaml
	kubectl wait --for=condition=complete pod -l app=urlshortener-test -n test --timeout=300s
	kubectl logs -n test -l app=urlshortener-test

upgrade-prod:
	helm repo add bitnami https://charts.bitnami.com/bitnami
	cd deployments/urlshortener && helm dependency build
	helm upgrade -n prod --wait --timeout 2m --values ./deployments/urlshortener/values.yaml urlshortener ./deployments/urlshortener

uninstall-test:
	kubectl delete -n test -f ./deployments/test/pod.yaml || true

clean:
	docker system prune -a -f

logs:
	kubectl wait -n test --timeout=300s --for=condition=complete job/urlshortener-test || kubectl logs -n test job/urlshortener-test; exit 1
	kubectl logs -n test job/urlshortener-test

default: get build
