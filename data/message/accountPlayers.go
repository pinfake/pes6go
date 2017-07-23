package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type AccountPlayers struct {
	Profiles []block.Piece
}

func (r AccountPlayers) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x3012, r.Profiles)...)

	return blocks
}

func NewAccountProfilesMessage(players block.AccountPlayers) AccountPlayers {
	return AccountPlayers{
		Profiles: block.GetPieces(reflect.ValueOf(players)),
	}
}
