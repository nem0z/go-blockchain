package blockchain

import (
	"bytes"
	"encoding/gob"
	"errors"

	badger "github.com/dgraph-io/badger/v3"

	"github.com/nem0z/go-blockchain/blockchain/block"
)

const DB_PATH = "./storage/blocks"

type Storage struct {
	DB *badger.DB
}

func InitDB() (error, *Storage) {
	badger, err := badger.Open(badger.DefaultOptions(DB_PATH))
	if err != nil {
		return err, &Storage{}
	}

	// db.DB = badger
	return nil, &Storage{DB: badger}
}

func (db *Storage) Load() (error, block.Hash) {
	var lastHash block.Hash

	err := db.DB.View(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			return errors.New("Not initialized blockchain")
		}

		if item, err := txn.Get([]byte("lh")); err != nil {
			return err
		} else {
			return item.Value(func(val []byte) error {
				if len(val) == len(lastHash) {
					copy(lastHash[:], val[:len(lastHash)])
					return nil
				}
				return errors.New("Expected [32]byte but lenght doesn't match with Hash lenght")
			})
		}
	})

	return err, lastHash
}

func (db *Storage) Add(b *block.Block) error {

	err := db.DB.Update(func(txn *badger.Txn) error {
		if err, serialized := Serialize(b); err != nil { // Serialize block
			return err
		} else {
			if err := txn.Set(b.Hash[:], serialized); err != nil { // Save block
				return err
			}

			return txn.Set([]byte("lh"), b.Hash[:]) // Save new block hash as LastHash
		}
	})

	return err
}

// Block to []byte and reverse

func Serialize(b *block.Block) (error, []byte) {
	var res bytes.Buffer

	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)

	return err, res.Bytes()
}

func Deserialize(data []byte) (error, *block.Block) {
	var block block.Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)

	return err, &block
}
