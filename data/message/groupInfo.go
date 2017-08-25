package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type GroupInfo struct {
	*block.GroupInfo
}

func (data GroupInfo) GetBlocks() []*block.Block {
	return block.GetBlocks(0x3052, data.GroupInfo)
}

func NewGroupInfoMessage(info *block.GroupInfo) GroupInfo {
	return GroupInfo{info}
}
