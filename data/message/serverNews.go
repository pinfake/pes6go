package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type ServerNews struct {
	News []*block.News
}

func (data ServerNews) GetBlocks() []*block.Block {
	var blocks []*block.Block
	blocks = append(blocks, block.GetBlocks(0x2009, block.Uint32{0})...)
	blocks = append(blocks, block.GetBlocks(0x200a, data.News)...)
	blocks = append(blocks, block.GetBlocks(0x200b, block.Uint32{0})...)
	return blocks
}

func NewServerNews(info []*block.News) ServerNews {
	return ServerNews{info}
}
