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

	fmt.Println("SpaceChain!")
	fmt.Println("Mining with difficulty", chain.MinBlockDifficulty)

	block.Mine(chain.MinBlockDifficulty)

	fmt.Println("Block", block)
	fmt.Println("Block's hash", block.Hash())

}
