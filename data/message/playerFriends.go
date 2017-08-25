package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type PlayerFriends struct {
	*block.PlayerFriends
}

func (PlayerFriends) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x3082, block.Uint32{0})...)
	blocks = append(blocks, block.GetBlocks(0x3086, block.Void{})...)
	return blocks
}

func NewPlayerFriendsMessage(info *block.PlayerFriends) PlayerFriends {
	return PlayerFriends{info}
}
