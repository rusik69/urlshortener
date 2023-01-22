package server

import "testing"

// generateKeyTest is a test for the generateKey function.
func TestGenerateKey(t *testing.T) {
	key := generateKey()
	if len(key) != 8 {
		t.Errorf("generateKey() = %s; want 6 characters", key)
	}
}
