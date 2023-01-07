package env

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Parse parses the environment variables and returnes the configuration.
func Parse() Config {
	listenPort := os.Getenv("SHORTENER_LISTEN_PORT")
	if listenPort == "" {
		listenPort = "8080"
	}
	listenHost := os.Getenv("SHORTENER_LISTEN_HOST")
	if listenHost == "" {
		listenHost = "0.0.0.0"
	}
	host := os.Getenv("SHORTENER_HOST")
	if host == "" {
		host = "http://localhost:8080"
	}
	dbHost := os.Getenv("SHORTENER_DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("SHORTENER_DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbUser := os.Getenv("SHORTENER_DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("SHORTENER_DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}
	dbName := os.Getenv("SHORTENER_DB_NAME")
	if dbName == "" {
		dbName = "postgres"
	}
	dbSSLMode := os.Getenv("SHORTENER_DB_SSLMODE")
	if dbSSLMode == "" {
		dbSSLMode = "disable"
	}
	dbTableName := os.Getenv("SHORTENER_DB_TABLE_NAME")
	if dbTableName == "" {
		dbTableName = "shortener"
	}

	logrus.Println("SHORTENER_LISTEN_PORT: ", listenPort)
	logrus.Println("SHORTENER_LISTEN_HOST: ", listenHost)
	logrus.Println("SHORTENER_HOST: ", host)
	logrus.Println("SHORTENER_DB_HOST: ", dbHost)
	logrus.Println("SHORTENER_DB_PORT: ", dbPort)
	logrus.Println("SHORTENER_DB_USER: ", dbUser)
	logrus.Println("SHORTENER_DB_PASSWORD: ", dbPassword)
	logrus.Println("SHORTENER_DB_NAME: ", dbName)
	logrus.Println("SHORTENER_DB_SSLMODE: ", dbSSLMode)
	logrus.Println("SHORTENER_DB_TABLE_NAME: ", dbTableName)

	return Config{
		ListenPort:  listenPort,
		ListenHost:  listenHost,
		Host:        host,
		DBHost:      dbHost,
		DBPort:      dbPort,
		DBUser:      dbUser,
		DBPassword:  dbPassword,
		DBName:      dbName,
		DBSSLMode:   dbSSLMode,
		DBTableName: dbTableName,
	}
}
