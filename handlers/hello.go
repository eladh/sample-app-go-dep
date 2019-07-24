package handlers

import (
	"fmt"
	"net/http"
)

// Handler for /
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s!\n\nWelcome to the Swisscom Application Cloud!", r.Header.Get("User-Agent"))
}
