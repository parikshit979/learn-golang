package user

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// User represents a user in the library
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// GenerateUserID generates a new user ID
func (u *User) GenerateUserID() {
	hash := sha256.New()
	hash.Write([]byte(u.FirstName + u.Email))
	u.ID = fmt.Sprintf("%x", string(hash.Sum(nil)))
}

// IsValidUser checks if the user is valid
func (u *User) IsValidUser() bool {
	if u.ID == "" {
		return false
	}

	if u.FirstName == "" {
		return false
	}

	if u.LastName == "" {
		return false
	}

	if u.Email == "" {
		return false
	}

	return true
}

// MarshalJSON marshals the user to JSON
func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	})
}

// UnmarshalJSON unmarshals the user from JSON
func (u *User) UnmarshalJSON(data []byte) error {
	type Alias User
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	return json.Unmarshal(data, aux)
}

// UsersStore represents a list of users
var UserStore = make(map[string]*User)

// GetUserByID returns the user by ID
func GetUserByID(id string) *User {
	return UserStore[id]
}

// AddUserToStore adds a new user
func AddUserToStore(user *User) {
	UserStore[user.ID] = user
}

// MarshalUsersJSON marshals the users to JSON
func MarshalUsersJSON() ([]byte, error) {
	return json.Marshal(UserStore)
}

// UnmarshalUsersJSON unmarshals the users from JSON
func UnmarshalUsersJSON(data []byte) error {
	return json.Unmarshal(data, &UserStore)
}
