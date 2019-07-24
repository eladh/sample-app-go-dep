package main

import (
	"github.com/nbari/violetear"
	"log"
	"net/http"
	"os"
	"sample-app-go-dep/handlers"
	"time"
)

var version string

func main() {
	version = os.Getenv("VERSION")

	router := violetear.New()
	router.HandleFunc("/health", health)
	router.HandleFunc("/", handlers.HelloHandler)

	srv := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   7 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(srv.ListenAndServe())
}
