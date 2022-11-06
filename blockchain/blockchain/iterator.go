package blockchain

import (
	badger "github.com/dgraph-io/badger/v3"
	"github.com/nem0z/go-blockchain/blockchain/block"
)

type BlockchainIterator struct {
	CurrentHash block.Hash
	Database    *Storage
}

func (bc *Blockchain) Iterator() (error, *BlockchainIterator) {
	err, lastHash := bc.LastHash()
	if err != nil {
		return err, &BlockchainIterator{}
	}

	iter := &BlockchainIterator{lastHash, bc.Database}
	return nil, iter
}

func (iter *BlockchainIterator) Next() (error, *block.Block) {
	var block *block.Block

	err := iter.Database.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash[:])
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			err, block = Deserialize(val)
			return err
		})

		return err
	})

	if err != nil {
		return err, block
	}

	iter.CurrentHash = block.PrevHash
	return err, block
}
