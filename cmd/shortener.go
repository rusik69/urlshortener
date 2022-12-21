package main

import (
	"github.com/rusik69/urlshortener/pkg/env"
	"github.com/rusik69/urlshortener/pkg/server"
)

func main() {
	config := env.Parse()
	
	server.Serve(config)
}
