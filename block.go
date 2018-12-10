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
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return block
}

// Create genesis block
func GenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte("0"))
}
