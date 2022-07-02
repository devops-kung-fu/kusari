package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"math"
	"math/big"
)

// Difficulty describes the difficulty of the algorithm
const Difficulty = 5

type Proof struct {
	Block  Block
	Target *big.Int
}

func (proof *Proof) init(nonce int) []byte {
	nonceHex, _ := ToHex(int64(nonce))
	difficultyHex, _ := ToHex(int64(Difficulty))
	data := bytes.Join(
		[][]byte{
			proof.Block.PrevHash,
			proof.Block.Data,
			nonceHex,
			difficultyHex,
		},
		[]byte{},
	)

	return data
}

// NewProof creates a Proof struct from a Block
// b is the Block object to use
func NewProof(b Block) (proof Proof) {
	target := big.NewInt(1)

	target.Lsh(target, uint(256-Difficulty))
	proof = Proof{b, target}

	return
}

func (proof *Proof) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		data := proof.init(nonce)
		hash = sha256.Sum256(data)

		log.Printf("\r%x", hash)

		intHash.SetBytes(hash[:])

		if intHash.Cmp(proof.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

func (proof *Proof) Validate() bool {
	var intHash big.Int

	data := proof.init(proof.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(proof.Target) == -1
}

func ToHex(num int64) (hex []byte, err error) {
	buff := new(bytes.Buffer)
	err = binary.Write(buff, binary.BigEndian, num)

	return buff.Bytes(), err
}
