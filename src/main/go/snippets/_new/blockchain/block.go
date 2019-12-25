package blockchain

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	Index        uint64
	Timestamp    string
	Transaction  *Transaction
	Hash         []byte
	PreviousHash []byte
}

func NewGenesisBlock() *Block {
	initialBlock := &Block{
		Index:       0,
		Timestamp:   time.Now().String(),
		Transaction: NewTransaction([]byte{}),
	}
	bHash, _ := GetTransactionHash(*initialBlock.Transaction)
	initialBlock.Hash = bHash
	initialBlock.PreviousHash = bHash
	return initialBlock
}

func GetBlockHash(b Block) ([]byte, error) {
	buf := bytes.NewBufferString("")
	buf.WriteString(string(b.Index))
	buf.WriteString(b.Timestamp)
	buf.Write(b.Transaction.Hash)
	buf.Write(b.PreviousHash)

	h := sha256.New()
	h.Write(buf.Bytes())

	return h.Sum(nil), nil
}
