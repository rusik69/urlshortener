package database

import (
	"database/sql"
	"fmt"
)

// InsertKey inserts the key url pair to database
func InsertKey(key, value string, db *sql.DB) error {
	sqlStmt := fmt.Sprintf("INSERT INTO %s(key, url) VALUES(%s, %s)", dbname, key, value)
	err := db.QueryRow(sqlStmt, key).Scan(&key)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		return nil
	}
	return nil
}
