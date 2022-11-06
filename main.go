package main

import (
	"log"

	"github.com/nem0z/go-blockchain/blockchain/blockchain"
)

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	err, bc := blockchain.New()
	Handle(err)

	// bc.CreateAndAdd("First block with data")

	err = bc.Display()
	Handle(err)
}
