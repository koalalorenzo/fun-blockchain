package chain

import "sort"

// BlockChain will allow us to sort blocks
type BlockChain []Block

func (b BlockChain) Len() int      { return len(b) }
func (b BlockChain) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b BlockChain) Less(i, j int) bool {
	return b[j].PreviousHash == b[i].HashString() && b[i].PreviousHash == "genesis"
}

// SortChain will ordinate the indexes of the blocks
func SortChain(blocks []Block) []Block {
	sort.Sort(BlockChain(blocks))
	return blocks
}
