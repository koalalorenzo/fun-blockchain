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
