package chain

import (
	"testing"
	"time"
)

func TestBlockDataExportImportHex(t *testing.T) {
	blockData := BlockData{
		time:      time.Now(),
		value:     "koalalorenzo",
		signature: "",
	}

	newHex := blockData.ToHex()
	newBlock := BlockDataFromHex(newHex)

	if !newBlock.time.Equal(blockData.time) {
		t.Error("Time failed")
	}
}
