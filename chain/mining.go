package chain

import "math/rand"

func miningThread(b *Block, difficulty float64) *Block {
	for {
		newNonce := rand.Uint32()
		b.Nonce = newNonce

		if b.IsHashValid(difficulty) {
			break
		}
	}

	return b
}

// Mine will try to find a valid random nonce
func (b *Block) Mine(difficulty float64) *Block {
	// Check if a valid nonce is already there
	if b.IsHashValid(difficulty) {
		return b
	}

	// Check if the content is valid
	if !b.IsContentValid() {
		return b
	}

	// Until the block has a valid hash, generate a new nonce
	b = miningThread(b, difficulty)
	return b
}
