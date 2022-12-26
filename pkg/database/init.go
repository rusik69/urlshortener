package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Init initializes the database
func Init(db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	sqlStmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (key TEXT PRIMARY KEY, url TEXT)", tableName)
	_, err := db.ExecContext(ctx, sqlStmt)
	if err != nil {
		return err
	}
	return nil
}
