package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type ServerNews struct {
	News []block.Piece
}

func (r ServerNews) GetBlocks() []*block.Block {
	var blocks []*block.Block
	blocks = append(blocks, block.GetBlocks(0x2009, []block.Piece{
		block.Uint32{0},
	})...)
	blocks = append(blocks, block.GetBlocks(0x200a, r.News)...)
	blocks = append(blocks, block.GetBlocks(0x200b, []block.Piece{
		block.Uint32{0},
	})...)
	return blocks
}

func NewServerNewsMessage(news []block.News) ServerNews {
	return ServerNews{
		News: block.GetPieces(reflect.ValueOf(news)),
	}
}
