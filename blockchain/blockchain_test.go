package blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlockChain(t *testing.T) {
	blockchain := NewBlockChain()
	assert.False(t, blockchain.IsEmpty())

	block, err := blockchain.Last()
	assert.Equal(t, "Genesis", string(block.Data))
	assert.NoError(t, err)
	assert.Empty(t, block.PrevHash)
}

func TestBlockChain_Serialize(t *testing.T) {
	blockchain := NewBlockChain()
	output, err := blockchain.Serialize()
	assert.NoError(t, err)
	assert.Len(t, output, 138)
}

func TestBlockChain_Pop(t *testing.T) {
	var blockchain BlockChain
	_, err := blockchain.Last()
	assert.Error(t, err)

	blockchain = NewBlockChain()
	block, err := blockchain.Last()
	assert.NoError(t, err)
	assert.Equal(t, []byte("Genesis"), block.Data)
}

func TestDeserialize(t *testing.T) {
	blockchain := NewBlockChain()
	output, _ := blockchain.Serialize()

	blockchain, err := Deserialize(output)
	assert.NoError(t, err)
	assert.Len(t, output, 138)
}

func TestBlockChain_Push(t *testing.T) {
	var blockchain BlockChain
	err := blockchain.Push([]byte("TEST"))
	assert.Error(t, err)

	blockchain = NewBlockChain()
	err = blockchain.Push([]byte("TEST"))
	assert.NoError(t, err)
	assert.Len(t, blockchain, 2)

	genesisBlock := blockchain[0]
	lastBlock, _ := blockchain.Last()

	assert.Equal(t, lastBlock.PrevHash, genesisBlock.Hash, "Previous hash of the last block in this test scenario should be the hash of the genesis block")
}
