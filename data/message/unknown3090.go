package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type Unknown3090 struct {
}

func (r Unknown3090) GetBlocks() []*block.Block {
	var blocks []*block.Block
	blocks = append(blocks, block.NewBlock(0x3091, block.Zero{}))
	return blocks
}

func NewUnknown3090Message() Unknown3090 {
	return Unknown3090{}
}
