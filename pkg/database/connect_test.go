package database

import (
	"math/rand"
	"testing"
	"time"

	"github.com/rusik69/urlshortener/pkg/env"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// generateKey generates a key.
func generateKey() string {
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// TestConnect tests the connection to the database
func TestDB(t *testing.T) {
	env.ConfigInstance.DBHost = "localhost"
	env.ConfigInstance.DBPort = "5432"
	env.ConfigInstance.DBPassword = "postgres"
	env.ConfigInstance.DBUser = "postgres"
	env.ConfigInstance.DBName = "urlshortener"
	env.ConfigInstance.DBSSLMode = "disable"
	env.ConfigInstance.DBTableName = "shortener"
	rand.Seed(time.Now().UTC().UnixNano())
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
	key := generateKey()
	t.Run("InsertKeyTest", func(t *testing.T) {
		err := InsertKey(key, "test")
		if err != nil {
			t.Errorf("InsertKey() = %s; want no error", err)
		}
	})
	t.Run("GetKeyTest", func(t *testing.T) {
		url, err := GetKey(key)
		t.Logf(key)
		t.Logf(url)
		if err != nil {
			t.Errorf("GetKey() = %s; want no error", err)
		}
		if url != "test" {
			t.Errorf("GetKey() = %s; want test", url)
		}
	})
}
