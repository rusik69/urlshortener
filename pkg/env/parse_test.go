package env

import "testing"

// TestParse tests the Parse function.
func TestParse(t *testing.T) {
	config := Parse()
	if config.listenPort != "8080" {
		t.Errorf("Expected listenPort to be 8080, got %s", config.listenPort)
	}
	if config.listenHost != "0.0.0.0" {
		t.Errorf("Expected listenHost to be 0.0.0.0, got %s", config.listenHost)
	}
	if config.dbHost != "localhost" {
		t.Errorf("Expected dbHost to be localhost, got %s", config.dbHost)
	}
	if config.dbPort != "5432" {
		t.Errorf("Expected dbPort to be 5432, got %s", config.dbPort)
	}
}
