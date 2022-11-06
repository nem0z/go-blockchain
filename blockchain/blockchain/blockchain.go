package blockchain

import (
	"fmt"

	"github.com/nem0z/go-blockchain/blockchain/block"
)

type Blockchain struct {
	Database *Storage
}

func New() (error, *Blockchain) {
	err, storage := InitDB()
	if err != nil {
		return err, &Blockchain{}
	}

	bc := &Blockchain{Database: storage}

	err, _ = bc.Database.Load()

	if err != nil {
		genesis := block.Genesis()
		if err := genesis.Validate(); err != nil {
			return err, &Blockchain{}
		}

		err := bc.Add(genesis)
		if err != nil {
			return err, &Blockchain{}
		}
	}

	return nil, bc
}

func (bc *Blockchain) Display() error {
	err, iter := bc.Iterator()

	if err != nil {
		return err
	}

	for {
		err, b := iter.Next()

		if err != nil {
			return err
		}

		fmt.Println("-----*-----")
		fmt.Printf("Hash => %x\n", b.Hash)
		fmt.Printf("Data => %s\n", b.Data)

		if b.PrevHash == (block.Hash{}) {
			break
		}
	}

	fmt.Println("-----*-----")
	return nil
}

func (bc *Blockchain) LastHash() (error, block.Hash) {
	return bc.Database.Load()
}

func (bc *Blockchain) Add(b *block.Block) error {

	if err := b.Validate(); err != nil {
		return err
	}

	bc.Database.Add(b)
	return nil
}

func (bc *Blockchain) CreateAndAdd(data string) error {
	err, lastHash := bc.LastHash()
	if err != nil {
		return err
	}

	newBlock := &block.Block{Data: []byte(data), PrevHash: lastHash}
	return bc.Add(newBlock)
}
