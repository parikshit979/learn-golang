package block

import (
	"encoding/json"
	"time"

	"github.com/learn-golang/blockchain/book"
)

// CurrentIndex represents the current index of the blockchain
var CurrentIndex int

// Block represents a block in the blockchain
type Block struct {
	Index       int                          `json:"index"`
	Timestamp   time.Time                    `json:"timestamp"`
	Transaction book.BookCheckoutTransaction `json:"transaction"`
	PrevHash    string                       `json:"prevHash"`
	Hash        string                       `json:"hash"`
}

// NewBlock creates a new Block instance
func NewBlock(transaction *book.BookCheckoutTransaction, prevHash string) *Block {
	timestamp := time.Now()
	hash := transaction.HashTransaction()

	return &Block{
		Index:       CurrentIndex + 1,
		Timestamp:   timestamp,
		Transaction: *transaction,
		PrevHash:    prevHash,
		Hash:        hash,
	}
}

// GenesisBlock creates the first block in the blockchain
func GenesisBlock(transaction *book.BookCheckoutTransaction) *Block {
	CurrentIndex = 0
	block := NewBlock(transaction, "")
	block.Index = CurrentIndex

	return block
}

// IsValidBlock checks if the block is valid
func (b *Block) IsValidBlock(prevBlock *Block) bool {
	if b.Index != prevBlock.Index+1 {
		return false
	}

	if b.PrevHash != prevBlock.Hash {
		return false
	}

	return true
}

// IsValidGenesisBlock checks if the block is a valid genesis block
func (b *Block) IsValidGenesisBlock() bool {
	if b.Index != 0 {
		return false
	}

	if b.PrevHash != "" {
		return false
	}

	return true
}

// MarshalJSON marshals the block to JSON
func (b *Block) MarshalJSON() ([]byte, error) {
	type Alias Block
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	})
}

// UnmarshalJSON unmarshals the block from JSON
func (b *Block) UnmarshalJSON(data []byte) error {
	type Alias Block
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	return json.Unmarshal(data, aux)
}
