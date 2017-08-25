package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type Unknown3120 struct {
}

func (r Unknown3120) GetBlocks() []*block.Block {
	var blocks []*block.Block
	blocks = append(blocks, block.GetBlocksFromPieces(0x3101, []block.Piece{
		block.Uint32{0},
	})...)
	blocks = append(blocks, block.GetBlocksFromPieces(0x3123, []block.Piece{
		block.Uint32{0},
	})...)
	return blocks
}

func NewUnknown3120Message() Unknown3120 {
	return Unknown3120{}
}
