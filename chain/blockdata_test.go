package chain

import (
	"testing"
	"time"
)

func TestBlockDataExportImportHex(t *testing.T) {
	blockData := BlockData{
		time:      time.Now(),
		value:     "koalalorenzo",
		signature: "proof",
	}

	newHex := blockData.ToHex()
	newBlock, _ := BlockDataFromHex(newHex)

	if newBlock.time.Unix() != newBlock.time.Unix() {
		t.Error("Time check failed")
	}

	if newBlock.value != "koalalorenzo" {
		t.Error("value check failed")
	}

	if newBlock.signature != "proof" {
		t.Error("Signature check failed")
	}
}

func TestBlockDataHashString(t *testing.T) {

	DataHex := "353437353635323034613735366332303332333632303332333033613331333033613332333532303433343535333534323033323330333133362c3662366636313663363136633666373236353665376136662c37303732366636663636"
	DataHash := "3811394768e7a30067d543e3ea81db73812af89cb6a1da3c4ae38df94029d185"

	blockData, err := BlockDataFromHex(DataHex)
	if err != nil {
		t.Error(err)
	}
	if blockData.HashString() != DataHash {
		t.Error("Hash is different")
	}
}
