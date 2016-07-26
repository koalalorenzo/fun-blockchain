package chain

import (
	"crypto"
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
	Cryptography
*/

// Sign will cryptographically sign the BlockData content (Hash of Value and Time)
func (data BlockData) Sign(privKey crypto.PrivateKey) {
	// ToDO
}

// IsSignatureValid will validate the signature of the block data
func (data BlockData) IsSignatureValid() bool {
	return false
}

/*
	Import / Export using other formats
*/

// ToHex returns the string hex of the BlockData values
func (data BlockData) ToHex() string {
	var sumArray []string

	// Encoding strings, since they may have commas
	unixDate := data.time.Format(time.UnixDate)
	timeHex := hex.EncodeToString([]byte(unixDate))
	sumArray = append(sumArray, timeHex)

	value := hex.EncodeToString([]byte(data.value))
	sumArray = append(sumArray, value)

	signature := hex.EncodeToString([]byte(data.signature))
	sumArray = append(sumArray, signature)

	// Dividing the hex values by comma
	sum := strings.Join(sumArray, ",")
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

// NewBlockData generates a new block data with a valid value
func NewBlockData(value string) BlockData {
	newBlockData := BlockData{
		time:  time.Now(),
		value: value,
	}
	return newBlockData
}
