package chain

import (
	"testing"
	"time"
)

func TestBlockExportImportHex(t *testing.T) {
	testBlock := Block{
		Time:         time.Now(),
		Nonce:        2,
		PreviousHash: "3",
	}

	testBlock.Data = append(testBlock.Data, BlockData{
		time:      time.Now(),
		value:     "woow",
		signature: "trust me",
	})

	newHex := testBlock.ToHex()
	newBlock := BlockFromHex(newHex)

	if newBlock.Time != testBlock.Time {
		t.Error("Time check failed")
	}

	if newBlock.Nonce != 2 {
		t.Error("Nonce check failed")
	}

	if newBlock.Data[0].signature != "trust me" {
		t.Error("Signature check failed")
	}

	if testBlock.IsHashValid(0) != newBlock.IsHashValid(0) {
		t.Error("Hash validation is different")
	}
}

func TestBlockHashString(t *testing.T) {

	hex := "313436393535363232313636383731383134362c322c332c333533343337333533363335333233303334363133373335333636333332333033333332333333363332333033333332333333303333363133333330333333333333363133333334333333313332333033343333333433353335333333353334333233303333333233333330333333313333333632633337333733363636333636363337333732633337333433373332333733353337333333373334333233303336363433363335"
	hash := "455c6c87cc43eb8477167fe9302bb531f30eb7927f2878c740bf63b07e59d7e9"

	block := BlockFromHex(hex)

	if block.HashString() != hash {
		t.Error("Hash is different")
	}
}

func TestBlockValidation(t *testing.T) {

	hex := "313436393535363232313636383731383134362c322c332c333533343337333533363335333233303334363133373335333636333332333033333332333333363332333033333332333333303333363133333330333333333333363133333334333333313332333033343333333433353335333333353334333233303333333233333330333333313333333632633337333733363636333636363337333732633337333433373332333733353337333333373334333233303336363433363335"

	block := BlockFromHex(hex)

	if block.IsContentValid() != true {
		t.Error("Block content not valid")
	}
}
