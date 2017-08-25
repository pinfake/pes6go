package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type ServerTime struct {
	ServerTime block.Piece
}

func (r ServerTime) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocksFromPieces(0x2007, []block.Piece{r.ServerTime})...)

	return blocks
}
