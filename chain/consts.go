package chain

/* Blocks Limits */

// MinBlockDifficulty defines a constant to consider a block valid.
// Check the Block.IsHashValid method
const MinBlockDifficulty = 1e+08

// MinBlockDataPerBlock defines the minimum data in a block.
// 0 = unused/empty blocks considered valid.
const MinBlockDataPerBlock = 0

// MaxBlockDataPerBlock defines the max lenght of the block's data array.
const MaxBlockDataPerBlock = 5

// MinSecsAfterBlock defines the seconds between a block and the next one
const MinSecsAfterBlock = 0
