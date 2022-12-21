package server

import "net/http"

// redirectHandler is the handler for the redirect path.
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
