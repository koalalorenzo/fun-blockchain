package chain

import (
	"crypto/sha256"
	"encoding/hex"
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

// Hash will return the hash of the entire block data.
func (data BlockData) Hash() []byte {
	hash := sha256.Sum256([]byte(data.ToHex()))
	doubleHash := sha256.Sum256(hash[:])
	return doubleHash[:]
}

// HashString returns the hash in string format
func (data BlockData) HashString() string {
	return hex.EncodeToString(data.Hash())
}

/*
	Import / Export using other formats
*/

// ToHex returns the string hex of the BlockData values
func (data BlockData) ToHex() string {

	// Encoding strings, since they may have commas
	unixDate := data.time.Format(time.UnixDate)
	timeHex := hex.EncodeToString([]byte(unixDate))

	value := hex.EncodeToString([]byte(data.value))
	signature := hex.EncodeToString([]byte(data.signature))

	// Dividing the hex values by comma
	sum := timeHex + "," + value + "," + signature
	return hex.EncodeToString([]byte(sum))
}

// BlockDataFromHex returns a new BlockData from a hex string
func BlockDataFromHex(hexString string) (BlockData, error) {

	blockHexs, _ := hex.DecodeString(hexString)
	blockArray := strings.Split(string(blockHexs), ",")

	newBlockData := BlockData{}

	timeString, err := hex.DecodeString(blockArray[0])
	if err != nil {
		panic(err)
	}
	timeValue, err := time.Parse(time.UnixDate, string(timeString))
	if err != nil {
		panic(err)
	}
	newBlockData.time = timeValue

	valueBytes, _ := hex.DecodeString(blockArray[1])
	newBlockData.value = string(valueBytes)

	signatureBytes, _ := hex.DecodeString(blockArray[2])
	newBlockData.signature = string(signatureBytes)

	if !newBlockData.IsValid() {
		panic("No valid block") // return error?
	}
	return newBlockData, nil
}
