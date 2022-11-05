package block

import (
	"bytes"
	"crypto/sha256"
)

type Hash [32]byte

type Block struct {
	Data     []byte
	PrevHash Hash
	Hash     Hash
}

func Genesis() *Block {
	return New("Genesis Block", Hash{})
}

func (b *Block) SetHash() {
	binBlock := bytes.Join([][]byte{b.PrevHash[:], b.Data[:]}, []byte{})
	hash := sha256.Sum256(binBlock)
	b.Hash = hash
}

func New(data string, prevHash Hash) *Block {
	block := &Block{Data: []byte(data), PrevHash: prevHash}
	block.SetHash()
	return block
}
