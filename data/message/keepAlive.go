package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type KeepAlive struct {
}

const keepAliveResponseQuery = 0x0005

func (m KeepAlive) GetBlocks() []block.Block {
	return []block.Block{
		block.NewBlock(keepAliveResponseQuery, block.Void{}),
	}
}
