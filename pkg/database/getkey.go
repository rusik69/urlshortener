package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rusik69/urlshortener/pkg/env"
)

// GetKey returns the key database
func GetKey(key string) (string, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	sqlStmt := fmt.Sprintf("SELECT url FROM %s WHERE key=$1;", env.ConfigInstance.DBTableName)
	err := DB.QueryRowContext(ctx, sqlStmt, key).Scan(&key)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", err
		}
		return "", nil
	}
	return key, nil
}
