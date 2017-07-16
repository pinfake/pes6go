package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type Unknown3100 struct {
}

func (r Unknown3100) GetBlocks() []block.Block {
	var blocks []block.Block
	blocks = append(blocks, block.NewBlock(0x3101, block.Zero{}))
	return blocks
}

func NewUnknown3100Message() Unknown3100 {
	return Unknown3100{}
}
