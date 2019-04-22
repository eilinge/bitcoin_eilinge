package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//Block :block struct
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int // 证明工作量
}

//NewBlock :genertor a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	// block.Hash = block.SetHash()
	// block.SetHash()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

//SetHash :genertor block hash by sha256
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	// hash [...]byte
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
	// return b.Hash
}

//NewgenesisBlock :create and returns genesis block
func NewgenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
