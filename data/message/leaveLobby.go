package message

import "github.com/pinfake/pes6go/data/block"

type LeaveLobby struct {
	Id uint32
}

func (data LeaveLobby) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4221, block.Uint32{data.Id})
}

func NewLeaveLobby(id uint32) LeaveLobby {
	return LeaveLobby{id}
}
