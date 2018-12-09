package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block struct
type Block struct {
	PrevBlockHash []byte
	Hash          []byte
	Data          []byte
	Timestamp     int64
}

// Set block hash value
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// Create new block
func NewBlock(data string, prevBlock []byte) *Block {
	block := &Block{prevBlock, []byte{}, []byte(data), time.Now().Unix()}
	block.SetHash()
	return block
}

// Create genesis block
func GenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte("0"))
}
