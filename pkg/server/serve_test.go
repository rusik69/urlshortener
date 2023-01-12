package server

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

// TestServe is a test for the Serve function.
func TestServe(t *testing.T) {
	host := "http://urlshortener:8080/"
	t.Run("RootTest", func(t *testing.T) {
		resp, err := http.Get(host)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if resp.StatusCode != 200 {
			t.Errorf("Expected status code to be 200, got %d", resp.StatusCode)
		}
	})
	var key string
	t.Run("ShortenTest", func(t *testing.T) {
		resp, err := http.Get(host + "shorten?url=" + url.QueryEscape("http://example.com/"))
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if resp.StatusCode != 200 {
			t.Errorf("Expected status code to be 200, got %d", resp.StatusCode)
		}
		keySplitBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		keySplit := strings.Split(string(keySplitBody), "/")
		key = keySplit[len(keySplit)-1]
		if len(key) != 6 {
			t.Logf("Key: %s\n", key)
			t.Errorf("Expected key to be 6 characters, got %d", len(key))
		}
	})
	t.Run("RedirectTest", func(t *testing.T) {
		url := host + key
		resp, err := http.Get(url)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code to be 301, got %d", resp.StatusCode)
		}
	})
}
