package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type Lobbies struct {
	pieces []block.Piece
}

func (r Lobbies) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x4201, r.pieces)...)

	return blocks
}

func NewLobbies(info block.Lobbies) Lobbies {
	return Lobbies{
		pieces: block.GetPieces(reflect.ValueOf(info)),
	}
}
