package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type Unknown3100 struct {
}

func (Unknown3100) GetBlocks() []*block.Block {
	return block.GetBlocks(0x3101, block.Uint32{0})
}

func NewUnknown3100() Unknown3100 {
	return Unknown3100{}
}
