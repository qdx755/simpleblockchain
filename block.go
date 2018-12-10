package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

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

// Serialize the block
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

// Deserialize the block
func Deserialize(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
