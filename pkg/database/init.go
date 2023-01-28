package database

import (
	"context"
	"fmt"
	"time"

	"github.com/rusik69/urlshortener/pkg/env"
	"github.com/sirupsen/logrus"
)

// Init initializes the database
func Init() error {
	sqlStmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (key TEXT PRIMARY KEY, url TEXT)", env.ConfigInstance.DBTableName)
	logrus.Println(sqlStmt)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	_, err := DB.ExecContext(ctx, sqlStmt)
	if err != nil {
		return err
	}
	return nil
}
