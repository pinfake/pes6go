package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type Unknown308c struct {
}

func (r Unknown308c) GetBlocks() []*block.Block {
	var blocks []*block.Block
	blocks = append(blocks, block.GetBlocks(0x308d, []block.Piece{
		block.Uint32{0},
	})...)
	return blocks
}

func NewUnknown308cMessage() Unknown308c {
	return Unknown308c{}
}
