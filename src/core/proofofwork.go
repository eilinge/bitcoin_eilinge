package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

/*
1. NewProofOfWork(block) (*ProofOfWork)
2. ProofOfWork{hash(block.nonce), target}.Run() (nonce int)
*/

var (
	maxNonce = math.MaxInt64
)

const targetBits = 20

// ProofOfWork ...
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork ...
func NewProofOfWork(b *Block) *ProofOfWork {
	// NewInt func(x int64) *Int
	target := big.NewInt(1)
	// Lsh sets z = x << n and returns z.
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{
		block:  b,
		target: target,
	}
	return pow
}

// perpareData :join struct and return []byte
func (pow *ProofOfWork) perpareData(nonce int) (data []byte) {
	// Join func(s [][]byte, sep []byte) []byte
	data = bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		IntToHex(pow.block.TimeStamp),
		// int64(int)
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce))},
		[]byte{})

	return
}

// Run :get nonce
func (pow *ProofOfWork) Run() ([]byte, int) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("\nMining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.perpareData(nonce)

		// crypto [Size]byte
		hash = sha256.Sum256(data)

		fmt.Printf("\r %x", hash)

		// SetBytes func(buf []byte) *Int
		hashInt.SetBytes(hash[:])

		// -1 if x < y /0 if x == y /+1 if x > y
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return hash[:], nonce
}

// Validate current block Nonce
func (pow *ProofOfWork) Validate() (isvalidate bool) {
	var hashInt big.Int
	data := pow.perpareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isvalidate = hashInt.Cmp(pow.target) == -1
	return
}
