package database

import (
	"context"
	"database/sql"
	"time"
)

// InsertKey inserts the key url pair to database
func InsertKey(key, value string, db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	sqlStmt := "INSERT INTO $1(key, url) VALUES($2, $3);"
	_, err := db.ExecContext(ctx, sqlStmt, tableName, key, value)
	if err != nil {
		return err
	}
	return nil
}
