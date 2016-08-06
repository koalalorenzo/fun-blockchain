package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

// Block structure
type Block struct {
	Time         time.Time
	Nonce        int64
	PreviousHash string
	Data         []BlockData
}

// Hash will return the hash of the entire block.
// Note that this method has to generate an unique value for the current block
func (b Block) Hash() []byte {
	hash := sha256.Sum256([]byte(b.ToHex()))
	doubleHash := sha256.Sum256(hash[:])
	return doubleHash[:]
}

// HashString returns the hash in string format
func (b Block) HashString() string {
	return hex.EncodeToString(b.Hash())
}

// IsContentValid Check if the Block's data is valid.
func (b Block) IsContentValid() bool {

	// Check if the block contains a minimum of data required.
	// if 0 then empty blocks are allowed.
	if len(b.Data) < MinBlockDataPerBlock {
		return false
	}

	// Check if the block has more elements than allowed
	// if 0 is considered infinite
	if MaxBlockDataPerBlock > 0 && len(b.Data) > MaxBlockDataPerBlock {
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
func (b Block) IsHashValid(difficulty int64) bool {
	if difficulty <= MinBlockDifficulty {
		difficulty = MinBlockDifficulty
	}

	magicHash := b.Hash()
	if int64(len(magicHash)) < difficulty {
		return false
	}
	// Checking if the bytes are 0 in the hash bytes
	for i := 0; i <= int(difficulty); i++ {
		if magicHash[i] != 0x0 {
			return false
		}
	}

	return true
}

/*
	Import / Export using other formats
*/

// ToHex returns the hexadecimal string of the Block
func (b Block) ToHex() string {
	var sum []string
	unixDate := b.Time.UnixNano()
	timeString := strconv.FormatInt(unixDate, 10)
	sum = append(sum, timeString)
	nonceString := strconv.FormatInt(b.Nonce, 10)
	sum = append(sum, nonceString)
	sum = append(sum, b.PreviousHash)

	if len(b.Data) > 0 {
		for _, nBlockData := range b.Data {
			sum = append(sum, nBlockData.ToHex())
		}
	}
	stringSum := strings.Join(sum, ",")
	return hex.EncodeToString([]byte(stringSum))
}

// BlockFromHex returns a new Block from a hex string
func BlockFromHex(hexString string) Block {

	blockHexs, _ := hex.DecodeString(hexString)
	blockArray := strings.Split(string(blockHexs), ",")
	newBlock := Block{}

	timeInt, err := strconv.ParseInt(blockArray[0], 10, 64)
	if err != nil {
		panic(err)
	}
	newBlock.Time = time.Unix(0, timeInt)

	nonceInt, err := strconv.ParseInt(blockArray[1], 10, 64)
	if err != nil {
		panic(err)
	}
	newBlock.Nonce = nonceInt

	newBlock.PreviousHash = blockArray[2]

	// Check if data is avaliable
	if len(blockArray) == 3 {
		return newBlock
	}

	for _, dataHex := range blockArray[3:] {
		dataBlock, _ := BlockDataFromHex(dataHex)
		// ToDo what if error?
		// ToDo Check if valid?
		newBlock.Data = append(newBlock.Data, dataBlock)
	}
	// ToDo Check if valid?
	return newBlock
}
