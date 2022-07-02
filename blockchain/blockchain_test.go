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

func TestBlockChain_Marshal(t *testing.T) {
	blockchain := NewBlockChain()
	output, err := blockchain.Marshal()
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

func TestUnMarshal(t *testing.T) {
	blockchain := NewBlockChain()
	output, _ := blockchain.Marshal()

	blockchain, err := UnMarshal(output)
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

func TestBlockChain_Iterator(t *testing.T) {
	blockchain := NewBlockChain()
	_ = blockchain.Push([]byte("1"))
	_ = blockchain.Push([]byte("2"))
	_ = blockchain.Push([]byte("3"))

	iterator := blockchain.Iterator()
	assert.Equal(t, len(*iterator.blockchain), iterator.current)

	c, _ := iterator.Next()
	assert.Equal(t, "3", string(c.Data))
	b, _ := iterator.Next()
	assert.Equal(t, "2", string(b.Data))
	assert.Equal(t, c.PrevHash, b.Hash)
	a, _ := iterator.Next()
	assert.Equal(t, "1", string(a.Data))
	assert.Equal(t, b.PrevHash, a.Hash)
	genesis, _ := iterator.Next()
	assert.Equal(t, a.PrevHash, genesis.Hash)
	assert.Equal(t, "Genesis", string(genesis.Data))
	_, err := iterator.Next()
	assert.Error(t, err)

}
