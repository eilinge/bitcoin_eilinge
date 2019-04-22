package main

import (
	"core"
	"fmt"
	"strconv"
)

func main() {
	bc := core.NewBlockchain()

	bc.AddBlock("send 1 BTC to Ivan")
	bc.AddBlock("send 2 more BTC to Ivan")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev.Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := core.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
