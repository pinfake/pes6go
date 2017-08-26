package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type Unknown3090 struct {
}

func (Unknown3090) GetBlocks() []*block.Block {
	return block.GetBlocks(0x3091, block.Uint32{0})
}

func NewUnknown3090() Unknown3090 {
	return Unknown3090{}
}
