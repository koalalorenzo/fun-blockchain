package main

import (
	"fmt"
	"time"

	"github.com/koalalorenzo/spacechain/chain"
)

func main() {
	block := chain.Block{
		Time: time.Now(),
	}
	fmt.Println("Mining with difficulty", chain.MinBlockDifficulty)

	block.Mine(chain.MinBlockDifficulty)
	fmt.Println("Time to mine it:", time.Since(block.Time))
	fmt.Println("Block's hash", block.Hash())
	fmt.Println("Genesis block created", block)

}
