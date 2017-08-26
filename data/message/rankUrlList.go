package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type RankUrlList struct {
	Urls []*block.RankUrl
}

func (data RankUrlList) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x2201, block.RankUrlsHeader{})...)
	blocks = append(blocks, block.GetBlocks(0x2202, data.Urls)...)
	blocks = append(blocks, block.GetBlocks(0x2203, block.Uint32{0})...)

	return blocks
}

func NewRankUrlListMessage(urls []*block.RankUrl) RankUrlList {
	return RankUrlList{urls}
}
