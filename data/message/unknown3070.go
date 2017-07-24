package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type Unknown3070 struct {
}

func (r Unknown3070) GetBlocks() []*block.Block {
	var blocks []*block.Block
	blocks = append(blocks, block.GetBlocks(0x3071, []block.Piece{
		block.Uint32{0},
	})...)
	blocks = append(blocks, block.GetBlocks(0x3073, []block.Piece{
		block.Uint32{0},
	})...)
	return blocks
}

func NewUnknown3070Message() Unknown3070 {
	return Unknown3070{}
}
