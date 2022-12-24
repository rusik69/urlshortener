package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// InsertKey inserts the key url pair to database
func InsertKey(key, value string, db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	sqlStmt := fmt.Sprintf("INSERT INTO %s(key, url) VALUES(%s, %s)", dbname, key, value)
	_, err := db.ExecContext(ctx, sqlStmt)
	if err != nil {
		return err
	}
	return nil
}
