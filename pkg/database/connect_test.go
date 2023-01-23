package database

import (
	"testing"

	"github.com/rusik69/urlshortener/pkg/env"
)

// TestConnect tests the connection to the database
func TestDB(t *testing.T) {
	env.ConfigInstance.DBHost = "postgres"
	env.ConfigInstance.DBPort = "5432"
	env.ConfigInstance.DBPassword = "postgres"
	env.ConfigInstance.DBUser = "postgres"
	env.ConfigInstance.DBName = "postgres"
	t.Run("ConnectTest", func(t *testing.T) {
		err := Connect()
		if err != nil {
			t.Errorf("Connect() = %s; want no error", err)
		}
	})
	t.Run("InitTest", func(t *testing.T) {
		err := Init()
		if err != nil {
			t.Errorf("Init() = %s; want no error", err)
		}
	})
	t.Run("InsertKeyTest", func(t *testing.T) {
		err := InsertKey("test", "test")
		if err != nil {
			t.Errorf("InsertKey() = %s; want no error", err)
		}
	})
	t.Run("GetKeyTest", func(t *testing.T) {
		url, err := GetKey("test")
		if err != nil {
			t.Errorf("GetKey() = %s; want no error", err)
		}
		if url != "test" {
			t.Errorf("GetKey() = %s; want test", url)
		}
	})
}
