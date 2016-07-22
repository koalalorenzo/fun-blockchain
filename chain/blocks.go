package chain

import (
	"encoding/hex"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/hashstructure"
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
	hashNumber := float64(b.Hash())

	remains := math.Mod(hashNumber, difficulty)

	if remains == 0 {
		return true
	}

	return false
}

/*
	Import / Export using other formats
*/

// ToHex returns the hexadecimal string of the Block
func (b Block) ToHex() string {

	var hexData []string
	for _, nBlockData := range b.Data {
		hexData = append(hexData, nBlockData.ToHex())
	}

	unixTime := b.Time.Unix()

	sum := string(unixTime) + "," + b.Nonce + "," + b.PreviousHash
	sum = sum + "," + strings.Join(hexData, ",")
	return hex.EncodeToString([]byte(sum))
}

// BlockFromHex returns a new Block from a hex string
func BlockFromHex(hexString string) Block {

	blockHexs, _ := hex.DecodeString(hexString)
	blockArray := strings.Split(string(blockHexs), ",")
	newBlock := Block{}

	unixTime, err := strconv.ParseInt(blockArray[0], 0, 64)
	if err != nil {
		panic(err)
	}
	newBlock.Time = time.Unix(unixTime, 0)

	newBlock.Nonce = blockArray[1]

	newBlock.PreviousHash = blockArray[2]

	for _, dataHex := range blockArray[3:] {
		dataBlock, _ := BlockDataFromHex(dataHex)
		// ToDo what if error?
		// ToDo Check if valid?
		newBlock.Data = append(newBlock.Data, dataBlock)
	}
	// ToDo Check if valid?
	return newBlock
}
