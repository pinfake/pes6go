package message

import "github.com/pinfake/pes6go/data/block"

const (
	initResponseQuery1 = 0x00002009 + iota
	initResponseQuery2
	initResponseQuery3
)

type Motd struct {
	Message block.ServerMessage
}

func (r Motd) GetBlocks() []block.Block {
	return []block.Block{
		block.NewBlock(initResponseQuery1, block.Zero{}),
		r.Message.GetBlock(initResponseQuery2),
		block.NewBlock(initResponseQuery3, block.Zero{}),
	}
}
