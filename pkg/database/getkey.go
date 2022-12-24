package database

import (
	"context"
	"database/sql"
	"time"
)

// GetKey returns the key database
func GetKey(key string, db *sql.DB) (string, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	sqlStmt := "SELECT url FROM shortener WHERE key=" + key + " LIMIT 1"
	err := db.QueryRowContext(ctx, sqlStmt, key).Scan(&key)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		}
		return "", nil
	}
	return key, nil
}
