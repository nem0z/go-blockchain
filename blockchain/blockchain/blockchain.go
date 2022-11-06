package blockchain

import (
	"fmt"
	"log"

	"github.com/nem0z/go-blockchain/blockchain/block"
)

type Blockchain struct {
	Blocks []*block.Block
}

func New() *Blockchain {
	genesis := block.Genesis()
	if err := genesis.Validate(); err != nil {
		log.Panic(err)
	}

	return &Blockchain{Blocks: []*block.Block{genesis}}
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

func (bc *Blockchain) AddBlock(b *block.Block) error {

	if err := b.Validate(); err != nil {
		return err
	}

	bc.Blocks = append(bc.Blocks, b)
	return nil
}

func (bc *Blockchain) CreateAndAdd(data string) error {
	newBlock := &block.Block{Data: []byte(data), PrevHash: bc.Last().PrevHash}

	return bc.AddBlock(newBlock)
}
