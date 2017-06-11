package message

import "github.com/pinfake/pes6go/data/block"

type RankUrlList struct {
	RankUrls block.RankUrls
}

func (r RankUrlList) GetBlocks() []block.Block {
	return []block.Block{
		block.NewBlock(0x2201, block.GenericBody{
			Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55},
		}),
		r.RankUrls.GetBlock(0x2202),
		block.NewBlock(0x2203, block.Zero{}),
	}
}
