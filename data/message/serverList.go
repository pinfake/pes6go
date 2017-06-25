package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type ServerList struct {
	Servers block.Servers
}

func (r ServerList) GetBlocks() []block.Block {
	var blocks []block.Block

	blocks = append(blocks, block.NewBlock(0x2002, block.Void{}))
	for _, bitBlock := range r.Servers.GetBlocks(0x2003) {
		blocks = append(blocks, bitBlock)
	}
	blocks = append(blocks, block.NewBlock(0x2004, block.Void{}))

	return blocks
}
