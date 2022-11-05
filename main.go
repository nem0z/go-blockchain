package main

import "github.com/nem0z/go-blockchain/blockchain/blockchain"

func main() {
	bc := blockchain.New()
	bc.CreateAndAddBlock("First block data")
	bc.CreateAndAddBlock("Second block data")
	bc.CreateAndAddBlock("Third block data")

	bc.Display()
}
