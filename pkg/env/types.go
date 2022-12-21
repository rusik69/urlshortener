package env

// Config is the configuration for the environment variables.
type Config struct {
	ListenPort string
	ListenHost string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DBSSLMode  string
}
