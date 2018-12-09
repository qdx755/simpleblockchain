package main

import "time"

// Block struct
type Block struct {
	PrevBlockHash []byte
	Hash          []byte
	Data          []byte
	Timestamp     int64
	Nonce         int
}

// Create new block
func NewBlock(data string, prevBlock []byte) *Block {
	block := &Block{prevBlock, []byte{}, []byte(data), time.Now().Unix(), 0}
	block.Mining()
	return block
}

// Create genesis block
func GenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte("0"))
}

// Mining to find the correct hash & nonce
func (b *Block) Mining() {
	nonce, hash := Prove(b)
	b.Hash = hash[:]
	b.Nonce = nonce
}
