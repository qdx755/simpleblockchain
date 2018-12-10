package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello Blockchain!")
	bc := NewBlockChain()
	bc.AddBlock("Send 1 BTC to Bob")
	bc.AddBlock("Send 2 BTC to Alice")
	for _, block := range bc.blocks {
		pow := NewProofOfWork(block)
		fmt.Printf("pow validate:%s\n", strconv.FormatBool(pow.validate()))
		fmt.Printf("%s with hash: %x\n", block.Data, block.Hash)
	}
}
