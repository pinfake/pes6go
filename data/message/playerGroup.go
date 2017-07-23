package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type PlayerGroup struct {
	PlayerGroup []block.Piece
}

func (r PlayerGroup) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x3042, r.PlayerGroup)...)

	return blocks
}

func NewPlayerGroupMessage(playerGroup block.PlayerGroup) PlayerGroup {
	return PlayerGroup{
		PlayerGroup: block.GetPieces(reflect.ValueOf(playerGroup)),
	}
}
