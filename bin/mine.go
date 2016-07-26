package main

import (
	"fmt"
	"time"

	"github.com/koalalorenzo/blockchain/chain"
)

func main() {
	genesisData := chain.NewBlockData("This is the unsigned genesis data")
	block := chain.Block{
		Time:         time.Now(),
		PreviousHash: "genesis",
		Data:         []chain.BlockData{genesisData},
	}
	fmt.Println("Mining with difficulty", chain.MinBlockDifficulty)

	block.Mine(chain.MinBlockDifficulty)
	fmt.Println("Time to mine it:", time.Since(block.Time))
	fmt.Println("Block's hash", block.HashString())
	fmt.Println("Block:", block)
	fmt.Println("Block created:")
	fmt.Println(block.ToHex())

	time.Sleep(60 * time.Second)

	secondData := chain.NewBlockData("This is the second block")
	secondBlock := chain.Block{
		Time:         time.Now(),
		PreviousHash: block.HashString(),
		Data:         []chain.BlockData{secondData},
	}

	secondBlock.Mine(chain.MinBlockDifficulty)
	fmt.Println("Time to mine 2 block:", time.Since(secondBlock.Time))
	fmt.Println("2 Block's hash", secondBlock.HashString())
	fmt.Println("2 Block:", secondBlock)
	fmt.Println("2 Block created:")
	fmt.Println(secondBlock.ToHex())
}
