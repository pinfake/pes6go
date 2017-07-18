package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type PlayerExtraSettings struct {
	PlayerExtraSettings []block.Piece
}

func (r PlayerExtraSettings) GetBlocks() []block.Block {
	var blocks []block.Block

	blocks = append(blocks, block.GetBlocks(0x4101, r.PlayerExtraSettings)...)

	return blocks
}

func NewPlayerExtraSettingsMessage(playerGroup block.PlayerExtraSettings) PlayerExtraSettings {
	return PlayerExtraSettings{
		PlayerExtraSettings: block.GetPieces(reflect.ValueOf(playerGroup)),
	}
}
