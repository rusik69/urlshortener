package database

import "database/sql"

// GetKey returns the key database
func GetKey(key string, db *sql.DB) (string, error) {
	sqlStmt := "SELECT url FROM shortener WHERE key=" + key + " LIMIT 1"
	err := db.QueryRow(sqlStmt, key).Scan(&key)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		}
		return "", nil
	}
	return key, nil
}
