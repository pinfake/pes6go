package common

import "github.com/pinfake/pes6go/network/blocks"

type KeepAlive struct {
}

const keepAliveResponseQuery = 0x0005

func (m KeepAlive) GetBlocks() []blocks.Block {
	return []blocks.Block{
		blocks.NewBlock(keepAliveResponseQuery, blocks.Void{}),
	}
}
