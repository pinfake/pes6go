package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

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

func NewMotdMessage(messages []block.ServerMessage) Motd {
	return Motd{
		Messages: block.GetPieces(reflect.ValueOf(messages)),
	}
}
