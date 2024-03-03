// blocks/blocks.go
package blocks

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Index     int
	Timestamp time.Time
	Data      string
	PrevHash  string
	Hash      string
}

// CalculateHash calculates the hash of a block
func (b *Block) CalculateHash() string {
	hashInput := fmt.Sprintf("%d%s%s%s", b.Index, b.Timestamp.String(), b.Data, b.PrevHash)
	hash := sha256.New()
	hash.Write([]byte(hashInput))
	return hex.EncodeToString(hash.Sum(nil))
}

// DisplayAllBlocks displays all blocks in the blockchain
func DisplayAllBlocks(blocks []Block) {
	for _, block := range blocks {
		DisplayBlock(block)
	}
}

// DisplayBlock displays information about a single block
func DisplayBlock(block Block) {
	fmt.Printf("Index: %d\n", block.Index)
	fmt.Printf("Timestamp: %s\n", block.Timestamp.String())
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Previous Hash: %s\n", block.PrevHash)
	fmt.Printf("Hash: %s\n", block.Hash)
	fmt.Println("------------------------------")
}

// NewBlock creates a new block in the blockchain
func NewBlock(index int, data string, prevHash string) Block {
	block := Block{
		Index:     index,
		Timestamp: time.Now(),
		Data:      data,
		PrevHash:  prevHash,
	}
	block.Hash = block.CalculateHash()
	return block
}

// ModifyBlock modifies the data of a block
func ModifyBlock(block *Block, newData string) {
	block.Data = newData
	block.Hash = block.CalculateHash()
	fmt.Println("Block modified successfully.")
}
