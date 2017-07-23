package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type GroupInfo struct {
	GroupInfo []block.Piece
}

func (r GroupInfo) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x3052, r.GroupInfo)...)

	return blocks
}

func NewGroupInfoMessage(groupInfo block.GroupInfo) GroupInfo {
	return GroupInfo{
		GroupInfo: block.GetPieces(reflect.ValueOf(groupInfo)),
	}
}
