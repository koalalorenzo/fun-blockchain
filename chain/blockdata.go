package chain

import (
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

// BlockData is the structure of the content of the block.
// Ideally any string value... consider this as a Transaction
type BlockData struct {
	time      time.Time
	value     string
	signature string
}

// IsValid check if the block's data is valid.
func (data BlockData) IsValid() bool {

	if data.time.After(time.Now()) {
		return false
	}

	return true
}

// ToHex returns the string hex of the BlockData values
func (data BlockData) ToHex() string {

	unixTime := data.time.Unix()
	value := hex.EncodeToString([]byte(data.value))
	signature := hex.EncodeToString([]byte(data.signature))

	// Dividing the hex values by comma
	sum := string(unixTime) + "," + value + "," + signature
	return hex.EncodeToString([]byte(sum))
}

// BlockDataFromHex returns a new BlockData from a hex string
func BlockDataFromHex(hexString string) BlockData {

	blockHexs, _ := hex.DecodeString(hexString)
	blockArray := strings.Split(string(blockHexs), ",")

	newBlock := BlockData{}

	unixTime, _ := strconv.ParseInt(blockArray[0], 0, 64)
	newBlock.time = time.Unix(unixTime, 0)

	valueBytes, _ := hex.DecodeString(blockArray[1])
	newBlock.value = string(valueBytes)

	signatureBytes, _ := hex.DecodeString(blockArray[2])
	newBlock.signature = string(signatureBytes)

	return newBlock
}
