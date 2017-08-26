package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type ServerList struct {
	Servers []*block.Server
}

func (data ServerList) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x2002, block.Void{})...)
	blocks = append(blocks, block.GetBlocks(0x2003, data.Servers)...)
	blocks = append(blocks, block.GetBlocks(0x2004, block.Void{})...)

	return blocks
}

func NewServerList(servers []*block.Server) ServerList {
	return ServerList{servers}
}
