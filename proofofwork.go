package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

// Mining difficulty hard code with const
const targetBits = 32

// Hash data with the block & nonce
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			intToHex(pow.block.Timestamp),
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
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(&pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}

func (pow *ProofOfWork) validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	validateFlag := hashInt.Cmp(&pow.target) == -1
	return validateFlag
}

type ProofOfWork struct {
	block  *Block
	target big.Int
}

// Create new proof of work
func NewProofOfWork(b *Block) *ProofOfWork {
	targetBytes := make([]byte, 32)
	target := big.Int{}
	// Make it easy to find the result, change the start number of zeros with 4
	numOfZeros := targetBits / 16
	targetBytes[numOfZeros-1] = 1
	target.SetBytes(targetBytes)
	pow := &ProofOfWork{b, target}
	return pow
}
