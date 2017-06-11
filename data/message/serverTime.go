package message

import (
	"github.com/pinfake/pes6go/data/block"
)

const (
	serverTimeResponseQuery1 = 0x00002007
)

type ServerTime struct {
	ServerTime block.ServerTime
}

func (r ServerTime) GetBlocks() []block.Block {
	return []block.Block{
		r.ServerTime.GetBlock(serverTimeResponseQuery1),
	}
}
