package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type RankUrlList struct {
	RankUrls []block.Piece
}

func (r RankUrlList) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.NewBlock(0x2201, block.GenericBody{
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55},
	}))
	blocks = append(blocks, block.GetBlocks(0x2202, r.RankUrls)...)
	blocks = append(blocks, block.GetBlocks(0x2203, []block.Piece{
		block.Uint32{0},
	})...)

	return blocks
}

func NewRankUrlListMessage(urls []block.RankUrl) RankUrlList {
	return RankUrlList{
		RankUrls: block.GetPieces(reflect.ValueOf(urls)),
	}
}
