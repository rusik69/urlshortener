package database

import (
	"context"
	"fmt"
	"time"

	"github.com/rusik69/urlshortener/pkg/env"
)

// InsertKey inserts the key url pair to database
func InsertKey(key, value string) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	sqlStmt := fmt.Sprintf("INSERT INTO %s(key, url) VALUES($1, $2);", env.ConfigInstance.DBTableName)
	_, err := DB.ExecContext(ctx, sqlStmt, key, value)
	if err != nil {
		return err
	}
	return nil
}
