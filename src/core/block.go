package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

/*
1. NewBlock(string, []byte) (Block)  // Block.Hash = []byte{}
2. Block.setHash()
1. NewGenesisBlock() (GenesisBlock)
*/

// Block struct
type Block struct {
	TimeStamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock ...
func NewBlock(data string, prevblockHash []byte) *Block {
	block := &Block{
		TimeStamp:     time.Now().Unix(), // int64
		Data:          []byte(data),
		PrevBlockHash: prevblockHash, // []byte
		Hash:          []byte{},
	}
	// block.SetHash()
	// get hash of current block
	pow := NewProofOfWork(block)
	block.Hash, block.Nonce = pow.Run()
	return block
}

// SetHash ...
func (b *Block) SetHash() []byte {
	// FormatInt func(i int64, base int) string
	timestamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
	// Join func(s [][]byte, sep []byte) []byte
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	// Sum256 func(data []byte) [Size]byte
	hash := sha256.Sum256(headers)
	// ([Size]byte[:]) []byte
	b.Hash = hash[:]
	return b.Hash
}

// NewGenesisBlock ...
func NewGenesisBlock() *Block {
	return NewBlock("create genesisBlock", []byte{})
}
