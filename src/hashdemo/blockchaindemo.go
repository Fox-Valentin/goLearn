package hashdemo

import "fmt"

// 定义区块链结构体
type BlockChain struct {
	Blocks []*Block
}

// 创建区块
func NewBlockChain() *BlockChain {
	bc := BlockChain{}
	genesisBlock := GenerateGenesisBlock()
	bc.AppendBlockChain(&genesisBlock)
	return &bc
}

// 验证
func isValid(bc *BlockChain, newBlock *Block) bool {
	preBlock := *bc.Blocks[len(bc.Blocks)-1]
	appendBlock := *newBlock
	flag := true
	if preBlock.Index != appendBlock.Index-1 {
		flag = false
	}
	if preBlock.Hash != appendBlock.PreBlockHash {
		flag = false
	}
	if appendBlock.Hash != calculateHash(*newBlock) {
		flag = false
	}
	return flag
}

// 上链方法
func (bc *BlockChain) AppendBlockChain(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else if isValid(bc, newBlock) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		fmt.Println("AppendBlockChain is invalid")
	}
}

//
func (bc *BlockChain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.AppendBlockChain(&newBlock)
}

// 查询
func (bc *BlockChain) Print() {
	for _, v := range bc.Blocks {
		fmt.Println(v)
	}
}
