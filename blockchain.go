package main

// Blockchain struct
type BlockChain struct {
	blocks []*Block
}

// Add block to blockchain
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// Create new blockchain
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{GenesisBlock()}}
}
