package chain

import (
	"math"
	"time"

	"github.com/mitchellh/hashstructure"
)

// Block structure
type Block struct {
	Time         time.Time
	Data         [MaxBlockDataPerBlock]BlockData
	Nonce        uint32
	PreviousHash uint64 // TBD
}

// Hash will return the hash of the entire block.
// Note that this method has to generate an unique value for the current block
func (b Block) Hash() uint64 {
	newHash, err := hashstructure.Hash(b, nil)

	if err != nil {
		panic(err)
	}

	return newHash
}

// IsContentValid Check if the Block's data is valid.
func (b Block) IsContentValid() bool {

	// Check if the block contains a minimum of data required.
	if len(b.Data) < MinBlockDataPerBlock {
		return false
	}

	// Validating the blockData
	for i := range b.Data {
		if !b.Data[i].IsValid() {
			return false
		}
	}

	return true
}

// IsHashValid check if the hash is valid with a specific difficulty.
// This method basically is defining when the algorythm has generated a valid
// block, using a specific difficulty.
// Alternatives of "algorythms" to validate it:
// 1. Trasnsform the hash in a string and check if it cointains "koalalorenzo"
// 2. Check if the hash contains by the answer to life universe and everything
func (b Block) IsHashValid(difficulty float64) bool {
	hashNumber := float64(b.Hash())

	remains := math.Mod(hashNumber, difficulty)

	if remains == 0 {
		return true
	}

	return false
}
