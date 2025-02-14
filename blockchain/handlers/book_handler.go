package handlers

import (
	"encoding/json"
	"net/http"

	blockImport "github.com/learn-golang/blockchain/block"
	bookImport "github.com/learn-golang/blockchain/book"
)

// AddBook adds a new book to the store
func AddBook(w http.ResponseWriter, r *http.Request) {
	// Parse the book from the request
	book := &bookImport.Book{}
	if err := json.NewDecoder(r.Body).Decode(book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate the book ID
	book.GenerateBookID()

	// Check if the book is valid
	if !book.IsValidBook() {
		http.Error(w, "Invalid book", http.StatusBadRequest)
		return
	}

	// Add the book
	bookImport.AddBookToStore(book)

	// Write the book
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Successfully added book"))
}

// GetBooks returns the list of books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Marshal the books
	data, err := bookImport.MarshalBooksJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the books
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// PurchaseBook purchases a book
func PurchaseBook(w http.ResponseWriter, r *http.Request) {
	// Parse the transaction from the request
	transaction := &bookImport.BookCheckoutTransaction{}
	if err := json.NewDecoder(r.Body).Decode(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the book exists
	book := bookImport.GetBookByID(transaction.BookID)
	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Check if the transaction is valid
	if !transaction.IsValidTransaction() {
		http.Error(w, "Invalid transaction", http.StatusBadRequest)
		return
	}

	var newBlock *blockImport.Block
	// Check if the transaction is a genesis transaction
	if transaction.IsGenesisTransaction() {
		blockImport.BlockChainStore = blockImport.NewBlockChain()
		newBlock = blockImport.GenesisBlock(transaction)
	} else {
		// Get the previous block
		prevBlock := blockImport.BlockChainStore.GetLatestBlock()

		// Create a new block
		newBlock = blockImport.NewBlock(transaction, prevBlock.Hash)

		// Check if the block is valid
		if !newBlock.IsValidBlock(prevBlock) {
			http.Error(w, "Invalid block", http.StatusBadRequest)
			return
		}
	}

	// Add the transaction to the blockchain
	blockImport.BlockChainStore.AddBlock(newBlock)

	// Write the transaction
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Successfully purchased book and added to blockchain"))
}
