package chain

import "time"

// BlockData is the structure of the content of the block.
// Ideally any string value... consider this as a Transaction
type BlockData struct {
	time  time.Time
	value string
}

// IsValid check if the block's data is valid.
func (data BlockData) IsValid() bool {

	if data.time.After(time.Now()) {
		return false
	}

	return true
}
