package core

// Blockchain ...
type Blockchain struct {
	Blocks []*Block
}

//AddBlock :add new block to blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

//NewBlockchain :genesis a new Blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewgenesisBlock()}}
}
