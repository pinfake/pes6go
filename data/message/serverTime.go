package message

import (
	"github.com/pinfake/pes6go/data/block"
)

const (
	serverTimeResponseQuery1 = 0x00002007
)

type ServerTime struct {
	ServerTime block.Piece
}

func (r ServerTime) GetBlocks() []block.Block {
	var blocks []block.Block

	for _, pieceBlock := range block.GetBlocks(serverTimeResponseQuery1, []block.Piece{r.ServerTime}) {
		blocks = append(blocks, pieceBlock)
	}
	return blocks
}
