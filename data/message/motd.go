package message

import "github.com/pinfake/pes6go/data/block"

type Motd struct {
	Messages []block.Piece
}

func (r Motd) GetBlocks() []block.Block {
	var blocks []block.Block
	blocks = append(blocks, block.NewBlock(0x2009, block.Zero{}))
	blocks = append(blocks, block.GetBlocks(0x200a, r.Messages)...)
	blocks = append(blocks, block.NewBlock(0x200b, block.Zero{}))

	return blocks
}
