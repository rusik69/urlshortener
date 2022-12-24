package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rusik69/urlshortener/pkg/env"
)

var dbname string

// Connect connects to the database.
func Connect(config env.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser,
		config.DBPassword, config.DBName, config.DBSSLMode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	dbname = config.DBName
	return db, nil
}
