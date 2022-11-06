package main

import "github.com/nem0z/go-blockchain/blockchain/blockchain"

func main() {
	bc := blockchain.New()
	bc.CreateAndAdd("First block data")
	bc.CreateAndAdd("Second block data")
	bc.CreateAndAdd("Third block data")

	bc.Display()
}
