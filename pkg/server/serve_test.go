package server

import (
	"net/http"
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
}
