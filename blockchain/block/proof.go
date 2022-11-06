package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

const Difficulty = 16

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func Proof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, (256 - Difficulty))
	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) Init(nonce int64) []byte {
	return bytes.Join(
		[][]byte{
			pow.Block.PrevHash[:],
			pow.Block.Data[:],
			ToHex(nonce),
			ToHex(Difficulty),
		},
		[]byte{},
	)
}

func (pow *ProofOfWork) Run() (int64, Hash) {
	var intHash big.Int
	var hash Hash
	var nonce int64

	nonce = 0

	for nonce < math.MaxInt64 {
		data := pow.Init(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("%x\r", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		}
		nonce++
	}

	fmt.Printf("\n")
	return nonce, hash
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.Init(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
