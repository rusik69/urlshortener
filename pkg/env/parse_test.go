package env

import "testing"

// TestParse tests the Parse function.
func TestParse(t *testing.T) {
	config := Parse()
	if config.ListenPort != "8080" {
		t.Errorf("Expected listenPort to be 8080, got %s", config.ListenPort)
	}
	if config.ListenHost != "0.0.0.0" {
		t.Errorf("Expected listenHost to be 0.0.0.0, got %s", config.ListenHost)
	}
	if config.Host != "http://localhost:8080" {
		t.Errorf("Expected host to be http://localhost:8080, got %s", config.Host)
	}
	if config.DBHost != "localhost" {
		t.Errorf("Expected dbHost to be localhost, got %s", config.DBHost)
	}
	if config.DBPort != "5432" {
		t.Errorf("Expected dbPort to be 5432, got %s", config.DBPort)
	}
	if config.DBUser != "postgres" {
		t.Errorf("Expected dbUser to be postgres, got %s", config.DBUser)
	}
	if config.DBPassword != "" {
		t.Errorf("Expected dbPassword to be empty, got %s", config.DBPassword)
	}
	if config.DBName != "postgres" {
		t.Errorf("Expected dbName to be postgres, got %s", config.DBName)
	}
	if config.DBSSLMode != "disable" {
		t.Errorf("Expected dbSSLMode to be disable, got %s", config.DBSSLMode)
	}
	if config.DBTableName != "shortener" {
		t.Errorf("Expected dbTableName to be shortener, got %s", config.DBTableName)
	}
}
