package env

import "os"

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
		dbPassword = ""
	}
	dbName := os.Getenv("SHORTENER_DB_NAME")
	if dbName == "" {
		dbName = "postgres"
	}
	dbSSLMode := os.Getenv("SHORTENER_DB_SSLMODE")
	if dbSSLMode == "" {
		dbSSLMode = "disable"
	}
	return Config{
		ListenPort: listenPort,
		ListenHost: listenHost,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbName:     dbName,
		DBSSLMode:  dbSSLMode,
	}
}
