package main

import (
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// Blockchain struct
type BlockChain struct {
	tip    []byte
}

// Iterator blockchain blocks
type BlockChainIterator struct {
	currentHash []byte
}

// Add block to blockchain
func (bc *BlockChain) AddBlock(data string) {
	//prevBlock := bc.blocks[len(bc.blocks)-1]
	//newBlock := NewBlock(data, prevBlock.Hash)
	//bc.blocks = append(bc.blocks, newBlock)
	var lastHash []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("1"))
		return nil
	})

	newBlock := NewBlock(data, lastHash)

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("1"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.tip = newBlock.Hash
		return nil
	})
}

// Iterator blockchain
func (bc *BlockChain) Iterator() *BlockChainIterator {
	bci := &BlockChainIterator{bc.tip}
	return bci
}

func (it *BlockChainIterator) Next() *Block {
	var block *Block
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(it.currentHash)
		block = Deserialize(encodedBlock)
		return nil
	})
	it.currentHash = block.PrevBlockHash
	return block
}

// Create new blockchain
func NewBlockChain() *BlockChain {
	//return &BlockChain{[]*Block{GenesisBlock()}}
	bc := BlockChain{}
	db,err:=bolt.Open(dbFile,0600,nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	err=db.Update(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(blocksBucket))
		if b == nil {
			genesisBlock := GenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}
			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			bc.tip = genesisBlock.Hash
		}else {
			tip := b.Get([]byte("1"))
			bc.tip = tip
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return &bc
}
