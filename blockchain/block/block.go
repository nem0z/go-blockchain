package block

import "errors"

type Hash [32]byte

type Block struct {
	Data     []byte
	Nonce    int64
	PrevHash Hash
	Hash     Hash
}

func New(data string, prevHash Hash) *Block {
	block := &Block{Data: []byte(data), PrevHash: prevHash}
	return block
}

func Genesis() *Block {
	return New("Genesis Block", Hash{})
}

func (b *Block) Validate() error {
	pow := Proof(b)
	nonce, hash := pow.Run()
	b.Nonce = nonce
	b.Hash = hash

	if pow.Validate() {
		return nil
	}
	return errors.New("Could not validate the block")
}
