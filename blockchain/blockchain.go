package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Hash [32]byte

type Blockchain struct {
	Blocks []*Block
}

type Block struct {
	Data     []byte
	PrevHash Hash
	Hash     Hash
}

// Block functions

func (b *Block) SetHash() {
	binBlock := bytes.Join([][]byte{b.PrevHash[:], b.Data[:]}, []byte{})
	hash := sha256.Sum256(binBlock)
	b.Hash = hash
}

func NewBlock(data string, prevHash Hash) *Block {
	block := &Block{Data: []byte(data), PrevHash: prevHash}
	block.SetHash()
	return block
}

// Blockchain functions

func NewBlockchain() *Blockchain {
	return &Blockchain{Blocks: []*Block{GenGenesisBlock()}}
}

func (bc *Blockchain) Display() {
	for _, b := range bc.Blocks {
		fmt.Println("-----*-----")
		fmt.Printf("Hash => %x\n", b.Hash)
		fmt.Printf("Data => %s\n", b.Data)
	}
	fmt.Println("-----*-----")
}

func GenGenesisBlock() *Block {
	return NewBlock("Genesis Block", Hash{})
}

func (bc *Blockchain) Last() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) AddBlock(b *Block) {
	b.SetHash()
	bc.Blocks = append(bc.Blocks, b)
}

func (bc *Blockchain) CreateAndAddBlock(data string) {
	newBlock := &Block{Data: []byte(data), PrevHash: bc.Last().PrevHash}
	bc.AddBlock(newBlock)
}
