package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type MenuServers struct {
	MenuServers []block.Piece
}

func (r MenuServers) GetBlocks() []block.Block {
	var blocks []block.Block

	blocks = append(blocks, block.GetBlocks(0x4201, r.MenuServers)...)

	return blocks
}

func NewMenuServersMessage(servers block.MenuServers) MenuServers {
	return MenuServers{
		MenuServers: block.GetPieces(reflect.ValueOf(servers)),
	}
}
