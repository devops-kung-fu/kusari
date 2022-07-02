package blockchain

import (
	"bytes"
	"encoding/gob"
	"errors"
	"log"
)

// Block defines a Block on the BlockChain.
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// BlockChain represents a BlockChain
type BlockChain []Block

// IsEmpty checks to see if the blockchain is empty
func (bc *BlockChain) IsEmpty() bool {
	return len(*bc) == 0
}

// Push pushes new data onto the blockchain
func (bc *BlockChain) Push(data []byte) (err error) {
	lastBlock, err := (*bc).Last()
	if err != nil {
		return
	}
	block, err := createBlock(data, lastBlock)
	*bc = append(*bc, block)
	return
}

// Pop removes and returns the last element of blockchain
func (bc *BlockChain) Last() (block Block, err error) {
	if bc.IsEmpty() {
		return block, errors.New("blockchain is empty")
	} else {
		index := len(*bc) - 1
		block = (*bc)[index]
		return block, nil
	}
}

// Serialize will serialize a BlockChain into an slice of bytes
func (bc *BlockChain) Serialize() (data []byte, err error) {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err = encoder.Encode(bc)
	return res.Bytes(), err
}

// Deserialize deserializes a slice of bytes into a BlockChain
// data the array of bytes to deserialize
func Deserialize(data []byte) (blockchain BlockChain, err error) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&blockchain)
	return
}

// NewBlockChain creates an empty BlockChain with a Genesis block
func NewBlockChain() (blockchain BlockChain) {
	block, _ := genesis()
	blockchain = append(blockchain, block)
	return
}

// CreateBlock Creates a new Block
// data is the string to be put in the block
// prevHash is the Has of the previous Block in the BlockChain
func createBlock(data []byte, lastBlock Block) (block Block, err error) {
	log.Println("Creating New Block...")
	block = Block{[]byte{}, data, lastBlock.Hash, 0}
	proof := NewProof(block)
	nonce, hash := proof.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return
}

// Genesis creates a genesis Block - the first Block in the BlockChain
func genesis() (block Block, err error) {
	var nullBlock Block
	return createBlock([]byte("Genesis"), nullBlock)
}
