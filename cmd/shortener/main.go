package main

import (
	"github.com/rusik69/urlshortener/pkg/database"
	"github.com/rusik69/urlshortener/pkg/env"
	"github.com/rusik69/urlshortener/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	env.Parse()
	logrus.SetReportCaller(true)
	err := database.Connect()
	if err != nil {
		panic(err)
	}
	err = database.Init()
	if err != nil {
		panic(err)
	}
	server.Serve()
}
