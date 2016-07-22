package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// Block structure
type Block struct {
	Time         time.Time
	Nonce        string
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
	if MaxBlockDataPerBlock > 0 && len(b.Data) < MaxBlockDataPerBlock {
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
	// magicHash := b.HashString()
	return true
}

/*
	Import / Export using other formats
*/

// ToHex returns the hexadecimal string of the Block
func (b Block) ToHex() string {

	unixDate := b.Time.Format(time.UnixDate)
	timeHex := hex.EncodeToString([]byte(unixDate))

	var hexData []string
	for _, nBlockData := range b.Data {
		hexData = append(hexData, nBlockData.ToHex())
	}

	sum := timeHex + "," + b.Nonce + "," + b.PreviousHash

	if len(hexData) > 0 {
		sum = sum + "," + strings.Join(hexData, ",")
	}

	fmt.Println("BlockSum", sum)
	return hex.EncodeToString([]byte(sum))
}

// BlockFromHex returns a new Block from a hex string
func BlockFromHex(hexString string) Block {

	blockHexs, _ := hex.DecodeString(hexString)
	blockArray := strings.Split(string(blockHexs), ",")
	newBlock := Block{}

	timeString, err := hex.DecodeString(blockArray[0])
	if err != nil {
		panic(err)
	}
	timeValue, err := time.Parse(time.UnixDate, string(timeString))
	if err != nil {
		panic(err)
	}
	newBlock.Time = timeValue

	newBlock.Nonce = blockArray[1]

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
