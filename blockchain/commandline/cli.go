package commandline

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/nem0z/go-blockchain/blockchain/blockchain"
)

type CommandLine struct {
	Blockchain *blockchain.Blockchain
}

func (cli *CommandLine) help() {
	fmt.Println("Usage :")
	fmt.Println(" add -block BLOCK_DATA - add a block to the chain")
	fmt.Println(" show - Prints the blocks in the chain")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.help()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	err := cli.Blockchain.CreateAndAdd(data)
	Handle(err)
	log.Println("Block added")
}

func (cli *CommandLine) showChain() {
	err := cli.Blockchain.Display()
	Handle(err)
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
		runtime.Goexit()
	}
}

func (cli *CommandLine) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	showChainCmd := flag.NewFlagSet("show", flag.ExitOnError)

	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		Handle(err)

	case "show":
		err := showChainCmd.Parse(os.Args[2:])
		Handle(err)

	default:
		cli.help()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if showChainCmd.Parsed() {
		cli.showChain()
	}
}
