package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type GamePlayerInfo struct {
	PlayerGroup []block.Piece
}

func (r GamePlayerInfo) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x4103, r.PlayerGroup)...)

	return blocks
}

func NewGamePlayerInfo(info block.PlayerExtended) GamePlayerInfo {
	return GamePlayerInfo{
		PlayerGroup: block.GetPieces(reflect.ValueOf(info)),
	}
}
