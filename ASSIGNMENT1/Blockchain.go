package main

import (
	"fmt"
	"crypto/sha256"
)

type block struct {
	nonce int
	data       string
	prevhash    string
	nextblock *block
}

func NewBlock(transaction string, nonce int, previousHash string) *block {
	return &block{nonce,transaction, previousHash, nil}
}

func ListBlocks(genesisBlock *block) {
	printBlock := genesisBlock
	for printBlock != nil {
		fmt.Printf("Transaction: %s\nNonce: %d\nPrevious Hash: %s\nCurrent Hash: %s\n\n", printBlock.data, printBlock.nonce, printBlock.prevhash, calculateHash(printBlock.data))
		printBlock = printBlock.nextblock
	}
}

func ChangeBlock(changeit *block,transaction string) {
	if changeit != nil {
		if changeit.data == transaction {
			changeit.data = "asharkhalil"
		}
	}
}


func VerifyChain(startingblock *block) bool {
	currentBlock := startingblock
	previousHash := ""

	for currentBlock != nil {
		currentHash := calculateHash(currentBlock.data)
		if currentBlock.prevhash != previousHash {
			fmt.Printf("Blockchain verification failed at block with data: %s\n", currentBlock.data)
			return false
		}
		if currentHash != calculateHash(currentBlock.data) {
			fmt.Printf("Block hash mismatch at block with data: %s\n", currentBlock.data)
			return false
		}

		previousHash = currentHash
		currentBlock = currentBlock.nextblock
	}

	fmt.Println("Blockchain verification passed. All blocks are valid.")
	return true
}

func calculateHash(data string)string{
	hash:=sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x",hash)
}

func main() {
	startBlock := NewBlock("ashar", 0, "")
	secondBlock := NewBlock("ali", 0, calculateHash(startBlock.data))
	thirdBlock := NewBlock("ahmed", 0, calculateHash(secondBlock.data))
	fourthBlock := NewBlock("amir", 0, calculateHash(thirdBlock.data))

	startBlock.nextblock = secondBlock
	secondBlock.nextblock = thirdBlock
	thirdBlock.nextblock = fourthBlock

	genesisBlock := startBlock

	ListBlocks(genesisBlock)
	ChangeBlock(secondBlock,"ali")
	VerifyChain(startBlock)
}
