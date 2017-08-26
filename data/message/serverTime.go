package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type ServerTime struct {
	*block.ServerTime
}

func (r ServerTime) GetBlocks() []*block.Block {
	return block.GetBlocks(0x2007, r.ServerTime)
}

func NewServerTime(info *block.ServerTime) ServerTime {
	return ServerTime{info}
}
