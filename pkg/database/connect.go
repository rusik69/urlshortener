package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/rusik69/urlshortener/pkg/env"
	"github.com/sirupsen/logrus"
)

// Connect connects to the database.
func Connect() error {
	var psqlInfo string
	if env.ConfigInstance.DBPassword == "" {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
			"sslmode=%s connect_timeout=60",
			env.ConfigInstance.DBHost, env.ConfigInstance.DBPort, env.ConfigInstance.DBUser,
			env.ConfigInstance.DBSSLMode)
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s sslmode=%s connect_timeout=60",
			env.ConfigInstance.DBHost, env.ConfigInstance.DBPort, env.ConfigInstance.DBUser,
			env.ConfigInstance.DBPassword, env.ConfigInstance.DBSSLMode)
	}
	logrus.Println(psqlInfo)
	success := false
	var db *sql.DB
	var err error
	for i := 1; i <= 10; i++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			logrus.Println("Error opening database connection, retrying")
			time.Sleep(5 * time.Second)
			continue
		}
		ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancelfunc()
		err = db.PingContext(ctx)
		if err != nil {
			logrus.Println("Error pinging database, retrying")
			time.Sleep(5 * time.Second)
			continue
		}
		success = true
	}
	if !success {
		fmt.Println("Reached max retries, exiting")
		return err
	}
	DB = db
	return nil
}
