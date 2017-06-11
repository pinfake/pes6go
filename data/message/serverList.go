package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type ServerList struct {
	Servers block.Servers
}

func (r ServerList) GetBlocks() []block.Block {
	return []block.Block{
		block.NewBlock(0x2002, block.Void{}),
		r.Servers.GetBlock(0x2003),
		block.NewBlock(0x2004, block.Void{}),
	}
}
