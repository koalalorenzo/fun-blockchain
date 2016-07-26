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

	if testBlock.IsHashValid(MinBlockDifficulty) != newBlock.IsHashValid(MinBlockDifficulty) {
		t.Error("Hash validation is different")
	}

	if testBlock.IsHashValid(MinBlockDifficulty) == true {
		t.Error("Block hash should be invalid")
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

	hex := "313436393535363937373332333032313136372c353537373030363739313934373737393431302c6e6f7065"
	hash := "739ed3d815b7af4b0e75e314adcb562f3a7b40651b0a993b0b9d01b98e0f7b01"
	block := BlockFromHex(hex)

	if block.HashString() != hash {
		t.Error("Block different")
	}

	if block.IsContentValid() != true {
		t.Error("Block content not valid")
	}

	if block.IsHashValid(MinBlockDifficulty) != true {
		t.Error("Block hash not acceptable")
	}
}

func TestBlockValidationData(t *testing.T) {
	hex := "313436393536323431313636363230343338342c353537373030363739313934373737393431302c6e6f70652c333533343337333533363335333233303334363133373335333636333332333033333332333333363332333033333332333333313333363133333334333333363333363133333335333333313332333033343333333433353335333333353334333233303333333233333330333333313333333632633335333433363338333633393337333333323330333633393337333333323330333733343336333833363335333233303337333533363635333733333336333933363337333636353336333533363334333233303336333733363335333636353336333533373333333633393337333333323330333633343336333133373334333633313263"
	hash := "b5078a59161143a7fe4a00d799fe847ce3235f39844c097c084c030ea9b55fa2"

	block := BlockFromHex(hex)

	if block.HashString() != hash {
		t.Error("Block different")
	}

	if len(block.Data) != 1 {
		t.Error("Block without data")
	}

	if block.Data[0].value != "This is the unsigned genesis data" {
		t.Error("Block Data has a different value")
	}

}
