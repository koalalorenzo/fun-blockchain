package chain

import "math/rand"

// Mine will try to find a valid random nonce
func (b *Block) Mine(difficulty int64) *Block {
	// Check if a valid nonce is already there
	if b.IsHashValid(difficulty) {
		return b
	}

	// Check if the content is valid
	if !b.IsContentValid() {
		return b
	}

	// Until the block has a valid hash, generate a new nonce
	// This is not using another method, nor a thread/goroutine, nor threads
	// because this code is designed to have fun
	for !b.IsHashValid(difficulty) {
		newNonce := rand.Int63()
		b.Nonce = newNonce
	}

	return b
}
