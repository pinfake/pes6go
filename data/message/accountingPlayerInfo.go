package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type AccountingPlayerInfo struct {
	PlayerExtended *block.PlayerExtended
}

func (data AccountingPlayerInfo) GetBlocks() []*block.Block {
	return block.GetBlocks(0x3042, data.PlayerExtended)
}

func NewAccountingPlayerInfo(info *block.PlayerExtended) AccountingPlayerInfo {
	return AccountingPlayerInfo{info}
}
