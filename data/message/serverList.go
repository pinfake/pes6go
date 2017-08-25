package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type ServerList struct {
	Servers []block.Piece
}

func (r ServerList) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocksFromPieces(0x2002, []block.Piece{
		block.Void{},
	})...)
	blocks = append(blocks, block.GetBlocksFromPieces(0x2003, r.Servers)...)
	blocks = append(blocks, block.GetBlocksFromPieces(0x2004, []block.Piece{
		block.Void{},
	})...)

	return blocks
}

func NewServerListMessage(servers []block.Server) ServerList {
	return ServerList{
		Servers: block.GetPieces(reflect.ValueOf(servers)),
	}
}
