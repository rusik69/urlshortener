package env

// Config is the configuration for the environment variables.
type Config struct {
	ListenPort   string
	ListenHost   string
	Host         string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBTableName  string
	DBSSLMode    string
	MaxURLLength int
}

// ConfigInstance is the instance of the Config struct.
var ConfigInstance Config
