package handlers

import (
	"encoding/json"
	"net/http"

	userImport "github.com/learn-golang/blockchain/user"
)

// AddUser adds a new user to the store
func AddUser(w http.ResponseWriter, r *http.Request) {
	// Parse the user from the request
	user := &userImport.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate the user ID
	user.GenerateUserID()

	// Check if the user is valid
	if !user.IsValidUser() {
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	}

	// Add the user
	userImport.AddUserToStore(user)

	// Write the user
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Successfully added user"))
}

// GetUsers returns the list of users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Marshal the users
	data, err := userImport.MarshalUsersJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the users
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
