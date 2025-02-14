package handlers

import (
	"net/http"

	blockImport "github.com/learn-golang/blockchain/block"
)

// GetBlockChain returns the blockchain
func GetBlockChain(w http.ResponseWriter, r *http.Request) {
	// Marshal the blockchain
	data, err := blockImport.BlockChainStore.MarshalBlockChainJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the blockchain
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
