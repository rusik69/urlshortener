FROM alpine

FROM --platform=$BUILDPLATFORM golang:1.18-alpine AS build-env
WORKDIR /go/src/github.com/rusik69/urlshortener
COPY . ./
ARG TARGETOS
ARG TARGETARCH
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -o /go/bin/shortener cmd/shortener/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /go/bin/shortener /app/shortener
COPY web /app/web

ENTRYPOINT ["/app/shortener"]
EXPOSE 8080