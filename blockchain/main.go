package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/learn-golang/blockchain/handlers"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Get the blockchain
	r.HandleFunc("/", handlers.GetBlockChain).Methods("GET")

	// Add a new book
	r.HandleFunc("/books", handlers.AddBook).Methods("POST")

	// Get the books
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")

	// Checkout a book
	r.HandleFunc("/checkout", handlers.PurchaseBook).Methods("POST")

	// Add a new user
	r.HandleFunc("/users", handlers.AddUser).Methods("POST")

	// Get the users
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")

	// Start the server
	log.Println("Listening on :3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
