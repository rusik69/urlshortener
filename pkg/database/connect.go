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

var tableName string

// Connect connects to the database.
func Connect(config env.Config) (*sql.DB, error) {
	var psqlInfo string
	if config.DBPassword == "" {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
			"dbname=%s sslmode=%s",
			config.DBHost, config.DBPort, config.DBUser,
			config.DBName, config.DBSSLMode)
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=%s",
			config.DBHost, config.DBPort, config.DBUser,
			config.DBPassword, config.DBName, config.DBSSLMode)
	}
	logrus.Println(psqlInfo)
	connectionFailed := false
	var db *sql.DB
	for i := 0; i < 10; i++ {
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			connectionFailed = true
			logrus.Println(err)
			time.Sleep(1 * time.Second)
			continue
		}
		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelfunc()
		err = db.PingContext(ctx)
		if err != nil {
			connectionFailed = true
			logrus.Println(err)
			time.Sleep(1 * time.Second)
			continue
		}
		if !connectionFailed {
			break
		}
	}
	tableName = config.DBTableName
	return db, nil
}
