package database

import (
	"context"
	"time"
)

// InsertKey inserts the key url pair to database
func InsertKey(key, value string) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	sqlStmt := "INSERT INTO shortener(key, url) VALUES($1, $2);"
	_, err := DB.ExecContext(ctx, sqlStmt, key, value)
	if err != nil {
		return err
	}
	return nil
}
