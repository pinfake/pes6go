package message

import "github.com/pinfake/pes6go/data/block"

type RankUrlList struct {
	RankUrls block.RankUrls
}

func (r RankUrlList) GetBlocks() []block.Block {
	var blocks []block.Block

	blocks = append(blocks, block.NewBlock(0x2201, block.GenericBody{
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55},
	}))
	for _, bitBlock := range r.RankUrls.GetBlocks(0x2202) {
		blocks = append(blocks, bitBlock)
	}
	blocks = append(blocks, block.NewBlock(0x2203, block.Zero{}))

	return blocks
}
