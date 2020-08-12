package hashdemo

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int64
	TimeStamp    int64
	PreBlockHash string
	Hash         string
	Data         string
}

func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.Data = data
	newBlock.PreBlockHash = preBlock.Hash
	newBlock.TimeStamp = time.Now().Unix()
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func GenerateGenesisBlock() Block {
	newBlock := Block{}
	newBlock.Index = -1
	newBlock.PreBlockHash = ""
	return GenerateNewBlock(newBlock, "Genesis Block")
}

func calculateHash(block Block) string {
	hashData := string(block.Index) + block.Data + block.PreBlockHash + string(block.TimeStamp)
	hash := sha256.Sum256([]byte(hashData))
	hashInHex := hex.EncodeToString(hash[:])
	return hashInHex
}
