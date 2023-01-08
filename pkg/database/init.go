package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// Init initializes the database
func Init(db *sql.DB) error {
	sqlStmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (key TEXT PRIMARY KEY, url TEXT)", dbConfig.DBTableName)
	logrus.Println(sqlStmt)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	logrus.Println(ctx)
	logrus.Println(db)
	logrus.Println(sqlStmt)
	res, err := db.ExecContext(ctx, sqlStmt)
	if err != nil {
		return err
	}
	logrus.Println(res)
	return nil
}
