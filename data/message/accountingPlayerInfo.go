package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type AccountingPlayerInfo struct {
	PlayerInfo []block.Piece
}

func (r AccountingPlayerInfo) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x3042, r.PlayerInfo)...)

	return blocks
}

func NewAccountingPlayerInfoMessage(info block.PlayerExtended) AccountingPlayerInfo {
	return AccountingPlayerInfo{
		PlayerInfo: block.GetPieces(reflect.ValueOf(info)),
	}
}
