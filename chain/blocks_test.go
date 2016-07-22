package chain

import (
	"testing"
	"time"
)

func TestBlockExportImportHex(t *testing.T) {
	testBlock := Block{
		Time:         time.Now(),
		Nonce:        "2",
		PreviousHash: "3",
	}

	testBlock.Data = append(testBlock.Data, BlockData{
		time:      time.Now(),
		value:     "woow",
		signature: "trust me",
	})

	newHex := testBlock.ToHex()
	t.Log("Econding done, Hex:", newHex)
	newBlock := BlockFromHex(newHex)
	t.Log("Decoding done", newBlock)

	if newBlock.Time != testBlock.Time {
		t.Error("Time check failed")
	}

	if newBlock.Nonce != "2" {
		t.Error("Nonce check failed")
	}

	if newBlock.Data[0].signature != "trust me" {
		t.Error("Signature check failed")
	}
}
