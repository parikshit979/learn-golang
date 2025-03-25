package main

import (
	"log"
	"net/http"

	"github.com/learn-golang/fileserverusinghttp/handlers"
)

func main() {
	http.HandleFunc("/get-url", handlers.GetURLHandler)
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("/path/to/files"))))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Failed to start server on port 3000")
	}
}
