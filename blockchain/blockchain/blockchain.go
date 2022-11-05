package blockchain

import (
	"fmt"

	"github.com/nem0z/go-blockchain/blockchain/block"
)

type Blockchain struct {
	Blocks []*block.Block
}

func New() *Blockchain {
	return &Blockchain{Blocks: []*block.Block{block.Genesis()}}
}

func (bc *Blockchain) Display() {
	for _, b := range bc.Blocks {
		fmt.Println("-----*-----")
		fmt.Printf("Hash => %x\n", b.Hash)
		fmt.Printf("Data => %s\n", b.Data)
	}
	fmt.Println("-----*-----")
}

func (bc *Blockchain) Last() *block.Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) AddBlock(b *block.Block) {
	b.SetHash()
	bc.Blocks = append(bc.Blocks, b)
}

func (bc *Blockchain) CreateAndAddBlock(data string) {
	newBlock := &block.Block{Data: []byte(data), PrevHash: bc.Last().PrevHash}
	bc.AddBlock(newBlock)
}
