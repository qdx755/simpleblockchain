package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

var (
	target   big.Int
	maxNonce = math.MaxInt64
)

// Mining difficulty hard code with const
const targetBits = 12

// Target difficulty, start with 6 zeros
// The smaller numOfZeros is, the easier to find the target
func setTarget() {
	targetBytes := make([]byte, 32)
	numOfZeros := targetBits / 4
	targetBytes[numOfZeros-1] = 1
	target.SetBytes(targetBytes)
}

// Hash data with the block & nonce
func prepareData(block *Block, nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			block.PrevBlockHash,
			block.Data,
			intToHex(block.Timestamp),
			intToHex(int64(targetBits)),
			intToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

// Covert int to hex
func intToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

// Find out the correct nonce & hash
func Prove(block *Block) (int, []byte) {
	setTarget()
	var hasInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining. the block data: %s\n", block.Data)
	for nonce < maxNonce {
		data := prepareData(block, nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("try nonce:%d with hash result: %x\n", nonce, hash)
		hasInt.SetBytes(hash[:])
		if hasInt.Cmp(&target) == -1 {
			fmt.Printf("find the block target nonce: %d\n", nonce)
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}
