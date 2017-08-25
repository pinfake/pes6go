package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type GamePlayerInfo struct {
	*block.PlayerExtended
}

func (data GamePlayerInfo) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4103, data.PlayerExtended)
}

func NewGamePlayerInfo(info *block.PlayerExtended) GamePlayerInfo {
	return GamePlayerInfo{info}
}
