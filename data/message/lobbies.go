package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type Lobbies struct {
	Lobbies []block.Piece
}

func (r Lobbies) GetBlocks() []block.Block {
	var blocks []block.Block

	blocks = append(blocks, block.GetBlocks(0x4201, r.Lobbies)...)

	return blocks
}

func NewLobbiesMessage(info block.Lobbies) Lobbies {
	return Lobbies{
		Lobbies: block.GetPieces(reflect.ValueOf(info)),
	}
}
