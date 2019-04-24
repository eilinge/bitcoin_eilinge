package core

// BlockChain ...
/*
1. NewBlockChain() (BlockChain)
2. BlockChain.AddBlock()
*/
type BlockChain struct {
	Blocks []*Block
}

// AddBlock ...
func (bc *BlockChain) AddBlock(data string) {
	preveBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, preveBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockChain ...
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
