package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type ServerList struct {
	Servers []block.Piece
}

func (r ServerList) GetBlocks() []block.Block {
	var blocks []block.Block

	blocks = append(blocks, block.NewBlock(0x2002, block.Void{}))
	blocks = append(blocks, block.GetBlocks(0x2003, r.Servers)...)
	blocks = append(blocks, block.NewBlock(0x2004, block.Void{}))

	return blocks
}
