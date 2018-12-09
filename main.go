package main

import "fmt"

func main() {
	fmt.Println("Hello Blockchain!")
	bc := NewBlockChain()
	bc.AddBlock("Send 1 BTC to Bob")
	bc.AddBlock("Send 2 BTC to Alice")
	for _, block := range bc.blocks {
		fmt.Printf("%s with hash: %x\n", block.Data, block.Hash)
	}
}
