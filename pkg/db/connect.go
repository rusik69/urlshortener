package db

import (
	"database/sql"
	"fmt"

	"github.com/rusik69/urlshortener/pkg/env"
)

// Connect connects to the database.
func Connect(config env.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)
	db, err := sql.Open("postgres", c.DSN())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
