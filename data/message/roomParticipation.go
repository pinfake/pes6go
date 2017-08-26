package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type RoomParticipation struct {
	*block.RoomParticipation
}

func (data RoomParticipation) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4365, data.RoomParticipation)
}

func NewRoomParticipation(info *block.RoomParticipation) RoomParticipation {
	return RoomParticipation{info}
}
