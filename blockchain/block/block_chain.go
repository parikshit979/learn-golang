package block

import "encoding/json"

// BlockChain represents the blockchain
type BlockChain struct {
	Blocks []*Block `json:"blocks"`
}

// BlockChainStore represents the blockchain
var BlockChainStore *BlockChain

// NewBlockChain creates a new BlockChain instance
func NewBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: make([]*Block, 0),
	}
}

// GenesisBlock creates the first block in the blockchain
func (bc *BlockChain) AddBlock(block *Block) {
	bc.Blocks = append(bc.Blocks, block)
}

// IsValidChain checks if the blockchain is valid
func (bc *BlockChain) IsValidChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		if !bc.Blocks[i].IsValidBlock(bc.Blocks[i-1]) {
			return false
		}
	}

	return true
}

// IsValidGenesisBlock checks if the blockchain is valid
func (bc *BlockChain) IsValidGenesisBlock() bool {
	return bc.Blocks[0].IsValidGenesisBlock()
}

// GetLatestBlock returns the latest block in the blockchain
func (bc *BlockChain) GetLatestBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

// GetBlockByIndex returns the block by index
func (bc *BlockChain) GetBlockByIndex(index int) *Block {
	return bc.Blocks[index]
}

// GetBlockByHash returns the block by hash
func (bc *BlockChain) GetBlockByHash(hash string) *Block {
	for _, block := range bc.Blocks {
		if block.Hash == hash {
			return block
		}
	}

	return nil
}

// GetBlockByPrevHash returns the block by previous hash
func (bc *BlockChain) GetBlockByPrevHash(prevHash string) *Block {
	for _, block := range bc.Blocks {
		if block.PrevHash == prevHash {
			return block
		}
	}

	return nil
}

// GetBlockByBookID returns the block by book ID
func (bc *BlockChain) GetBlockByBookID(bookID string) *Block {

	for _, block := range bc.Blocks {
		if block.Transaction.BookID == bookID {
			return block
		}
	}

	return nil
}

// GetBlockByUserID returns the blocks by user ID
func (bc *BlockChain) GetBlockByUserID(userID string) *Block {

	for _, block := range bc.Blocks {
		if block.Transaction.UserID == userID {
			return block
		}

	}

	return nil
}

// GetBlockByCheckoutDate returns the blocks by checkout date
func (bc *BlockChain) GetBlockByCheckoutDate(checkoutDate string) *Block {
	for _, block := range bc.Blocks {
		if block.Transaction.CheckoutDate == checkoutDate {
			return block
		}
	}

	return nil
}

// MarshalBlockChainJSON marshals the blockchain to JSON
func (bc *BlockChain) MarshalBlockChainJSON() ([]byte, error) {
	type Alias BlockChain
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(bc),
	})
}

// UnmarshalBlockChainJSON unmarshals the blockchain from JSON
func (bc *BlockChain) UnmarshalBlockChainJSON(data []byte) error {
	type Alias BlockChain
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(bc),
	}
	return json.Unmarshal(data, aux)
}
