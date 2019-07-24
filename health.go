package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// health returns: status code 200 and  time & version in json format
func health(w http.ResponseWriter, r *http.Request) {
	j := map[string]string{
		"time":    time.Now().UTC().Format(time.RFC3339),
		"version": version,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(j); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
