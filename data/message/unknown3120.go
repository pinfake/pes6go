package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type Unknown3120 struct {
}

func (Unknown3120) GetBlocks() []*block.Block {
	var blocks []*block.Block
	blocks = append(blocks, block.GetBlocks(0x3101, block.Uint32{0})...)
	blocks = append(blocks, block.GetBlocks(0x3123, block.Uint32{0})...)
	return blocks
}

func NewUnknown3120() Unknown3120 {
	return Unknown3120{}
}
