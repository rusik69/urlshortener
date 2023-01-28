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
	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	sqlStmt := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", env.ConfigInstance.DBName)
	logrus.Println(sqlStmt)
	sqlStmt = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (key TEXT PRIMARY KEY, url TEXT)", env.ConfigInstance.DBTableName)
	logrus.Println(sqlStmt)
	_, err := DB.ExecContext(ctx, sqlStmt)
	if err != nil {
		return err
	}
	sqlStmt = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (key TEXT PRIMARY KEY, url TEXT)", env.ConfigInstance.DBTableName)
	logrus.Println(sqlStmt)
	_, err = DB.ExecContext(ctx, sqlStmt)
	if err != nil {
		return err
	}
	return nil
}
