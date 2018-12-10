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
	bci := bc.Iterator()
	for {
		block := bci.Next()
		pow := NewProofOfWork(block)
		fmt.Printf("pow validate:%s\n", strconv.FormatBool(pow.validate()))
		fmt.Printf("%s with hash: %x\n", block.Data, block.Hash)
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
