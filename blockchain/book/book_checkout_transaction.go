package book

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// BookCheckoutTransaction represents a book checkout transaction details
type BookCheckoutTransaction struct {
	BookID       string `json:"bookID"`
	UserID       string `json:"userID"`
	IsGenesis    bool   `json:"isGenesis"`
	CheckoutDate string `json:"checkoutDate"`
}

// IsGenesisTransaction checks if the transaction is a genesis transaction
func (t *BookCheckoutTransaction) IsGenesisTransaction() bool {
	return t.IsGenesis
}

// IsValidTransaction checks if the transaction is valid
func (t *BookCheckoutTransaction) IsValidTransaction() bool {
	if t.BookID == "" {
		return false
	}

	if t.UserID == "" {
		return false
	}

	if t.CheckoutDate == "" {
		return false
	}

	return true
}

// HashTransaction generates a hash for the transaction
func (t *BookCheckoutTransaction) HashTransaction() string {
	hash := sha256.New()
	hash.Write([]byte(t.BookID + t.UserID + t.CheckoutDate))
	return fmt.Sprintf("%x", string(hash.Sum(nil)))
}

// MarshalTransactionJSON marshals the transaction to JSON
func (t *BookCheckoutTransaction) MarshalTransactionJSON() ([]byte, error) {
	type Alias BookCheckoutTransaction
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(t),
	})
}

// UnmarshalTransactionJSON unmarshals the transaction from JSON
func (t *BookCheckoutTransaction) UnmarshalTransactionJSON(data []byte) error {
	type Alias BookCheckoutTransaction
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	return json.Unmarshal(data, aux)
}
