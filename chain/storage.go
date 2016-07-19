package chain

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/structs"
)

// ToMap will return a block into a map
func (b Block) ToMap() map[string]interface{} {
	return structs.Map(b)
}

// ToJSON will return a block struct into json
func (b *Block) ToJSON() []byte {
	blockJSON, err := json.Marshal(b.ToMap())
	if err != nil {
		fmt.Println("Error:", err)
	}
	os.Stdout.Write(blockJSON)
	return blockJSON
}
