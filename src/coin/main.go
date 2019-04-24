package main

import (
	"core"
	"fmt"
)

func main() {
	bc := core.NewBlockChain()

	bc.AddBlock("bobe send 1 bitcoin to eilin")
	bc.AddBlock("eilin send 1 bitcoin to duzi")

	for _, block := range bc.Blocks {
		fmt.Println("")
		fmt.Printf("get block.Hash: %x\n", block.Hash)
		fmt.Printf("get block.Data: %v\n", string(block.Data))
		fmt.Printf("get block.PrevBlockHash: %x\n", block.PrevBlockHash)

		pow := core.NewProofOfWork(block)
		isvalidate := pow.Validate()
		fmt.Printf("validate the block nonce is: %v\n", isvalidate)
		fmt.Println("")
	}
}
