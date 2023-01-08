FROM alpine

FROM golang:1.18-alpine AS build-env
WORKDIR /go/src/github.com/rusik69/urlshortener
COPY . ./
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
    go build -o /go/bin/shortener cmd/shortener/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /go/bin/shortener /app/shortener
COPY web /app/web

ENTRYPOINT ["/app/shortener"]
EXPOSE 8080