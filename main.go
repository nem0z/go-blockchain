package main

import (
	"log"
	"os"

	"github.com/nem0z/go-blockchain/blockchain/blockchain"
	cl "github.com/nem0z/go-blockchain/blockchain/commandline"
)

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	defer os.Exit(0)
	err, bc := blockchain.New()
	Handle(err)

	defer bc.Close()

	// bc.CreateAndAdd("First block with data")

	cli := cl.CommandLine{Blockchain: bc}
	cli.Run()
}
