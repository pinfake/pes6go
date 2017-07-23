package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type Unknown3120 struct {
}

func (r Unknown3120) GetBlocks() []*block.Block {
	var blocks []*block.Block
	blocks = append(blocks, block.NewBlock(0x3101, block.Zero{}))
	blocks = append(blocks, block.NewBlock(0x3123, block.Zero{}))
	return blocks
}

func NewUnknown3120Message() Unknown3120 {
	return Unknown3120{}
}
