package main

import (
	"github.com/rusik69/urlshortener/pkg/database"
	"github.com/rusik69/urlshortener/pkg/env"
	"github.com/rusik69/urlshortener/pkg/server"
)

func main() {
	config := env.Parse()
	db, err := database.Connect(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = database.Init(db)
	if err != nil {
		panic(err)
	}
	server.Serve(config, db)
}
