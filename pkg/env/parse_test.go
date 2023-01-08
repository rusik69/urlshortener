package env

import (
	"testing"
)

// TestParse tests the Parse function.
func TestParse(t *testing.T) {
	if ConfigInstance.ListenPort != "8080" {
		t.Errorf("Expected listenPort to be 8080, got %s", ConfigInstance.ListenPort)
	}
	if ConfigInstance.ListenHost != "0.0.0.0" {
		t.Errorf("Expected listenHost to be 0.0.0.0, got %s", ConfigInstance.ListenHost)
	}
	if ConfigInstance.Host != "http://localhost:8080" {
		t.Errorf("Expected host to be http://localhost:8080, got %s", ConfigInstance.Host)
	}
	if ConfigInstance.DBHost != "localhost" {
		t.Errorf("Expected dbHost to be localhost, got %s", ConfigInstance.DBHost)
	}
	if ConfigInstance.DBPort != "5432" {
		t.Errorf("Expected dbPort to be 5432, got %s", ConfigInstance.DBPort)
	}
	if ConfigInstance.DBUser != "postgres" {
		t.Errorf("Expected dbUser to be postgres, got %s", ConfigInstance.DBUser)
	}
	if ConfigInstance.DBPassword != "postgres" {
		t.Errorf("Expected dbPassword to be postgres, got %s", ConfigInstance.DBPassword)
	}
	if ConfigInstance.DBName != "postgres" {
		t.Errorf("Expected dbName to be postgres, got %s", ConfigInstance.DBName)
	}
	if ConfigInstance.DBSSLMode != "disable" {
		t.Errorf("Expected dbSSLMode to be disable, got %s", ConfigInstance.DBSSLMode)
	}
	if ConfigInstance.DBTableName != "shortener" {
		t.Errorf("Expected dbTableName to be shortener, got %s", ConfigInstance.DBTableName)
	}
}
