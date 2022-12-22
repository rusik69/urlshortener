package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rusik69/urlshortener/pkg/env"
)

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
		panic(err)
	}
	return db, nil
}
