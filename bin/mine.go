package main

import (
	"fmt"
	"time"

	"github.com/koalalorenzo/blockchain/chain"
)

func main() {
	block := chain.Block{
		Time:         time.Now(),
		PreviousHash: "nope",
	}
	fmt.Println("Mining with difficulty", chain.MinBlockDifficulty)

	block.Mine(chain.MinBlockDifficulty)
	fmt.Println("Time to mine it:", time.Since(block.Time))
	fmt.Println("Block's hash", block.Hash())
	fmt.Println("Block:", block)
	fmt.Println("Genesis block created:")
	fmt.Println(block.ToHex())
	fmt.Println(chain.BlockFromHex(block.ToHex()))
}
