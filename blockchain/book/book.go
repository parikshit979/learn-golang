package book

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// Book represents a book in the library
type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	PublishDate string `json:"publishDate"`
}

// GenerateBookID generates a new book ID
func (b *Book) GenerateBookID() {
	hash := sha256.New()
	hash.Write([]byte(b.ISBN + b.PublishDate))
	b.ID = fmt.Sprintf("%x", string(hash.Sum(nil)))
}

// IsValidBook checks if the book is valid
func (b *Book) IsValidBook() bool {
	if b.ID == "" {
		return false
	}

	if b.Title == "" {
		return false
	}

	if b.Author == "" {
		return false
	}

	if b.ISBN == "" {
		return false
	}

	if b.PublishDate == "" {
		return false
	}

	return true
}

// MarshalJSON marshals the book to JSON
func (b *Book) MarshalJSON() ([]byte, error) {
	type Alias Book
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	})
}

// UnmarshalJSON unmarshals the book from JSON
func (b *Book) UnmarshalJSON(data []byte) error {
	type Alias Book
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	return json.Unmarshal(data, aux)
}

// Books represents a list of books
var BookStore = make(map[string]*Book)

// GetBookByID returns the book by ID
func GetBookByID(id string) *Book {
	return BookStore[id]
}

// AddBookToStore adds a new book
func AddBookToStore(book *Book) {
	BookStore[book.ID] = book
}

// MarshalBooks marshals the books to JSON
func MarshalBooksJSON() ([]byte, error) {
	return json.Marshal(BookStore)
}

// UnmarshalBooks unmarshals the books from JSON
func UnmarshalBooksJSON(data []byte) error {
	return json.Unmarshal(data, &BookStore)
}
