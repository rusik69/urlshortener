package server

import (
	"database/sql"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
maxAttempts = 10

// randSeq generates a random string of the given length.
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// KeyExists checks if the given key exists in the database.
func KeyExists(db *sql.DB, key string) (bool, error) {
	sqlStmt := "SELECT url FROM shortener WHERE key=" + key + " LIMIT 1"
	err := db.QueryRow(sqlStmt, key).Scan(&key)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

// generateKey generates a key for the given URL.
func generateKey(url string) (string, error) {
	key := randSeq(6)
	attempt := 1
	for keyExists, err := KeyExists(DB, key); keyExists; {
		if err != nil || attempt > maxAttempts {
			return "", err
		}
		key = randSeq(6)
	}
	return key, nil
}
