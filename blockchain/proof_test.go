package blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProof_Validate(t *testing.T) {
	blockchain := NewBlockChain()
	_ = blockchain.Push([]byte("1"))

	iterator := blockchain.Iterator()
	assert.Equal(t, len(*iterator.blockchain), iterator.current)

	block, _ := iterator.Next()

	proof := NewProof(block)
	valid := proof.Validate()

	assert.True(t, valid)
}
