package server

import "github.com/rusik69/urlshortener/pkg/database"

// saveURL saves the key, URL pair to the database.
func saveURL(url, key string) error {
	return database.InsertKey(key, url)
}
