package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type PlayerUpdate struct {
	Player []block.Piece
}

func (r PlayerUpdate) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x4220, r.Player)...)

	return blocks
}

func NewPlayerUpdateMessage(player block.Player) PlayerUpdate {
	return PlayerUpdate{
		Player: block.GetPieces(reflect.ValueOf(player)),
	}
}
