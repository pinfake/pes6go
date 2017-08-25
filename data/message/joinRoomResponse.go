package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type JoinRoomResponse struct {
	*block.JoinRoomResponse
}

func (data JoinRoomResponse) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4321, data.JoinRoomResponse)
}

func NewJoinRoomResponse(code uint32, position byte) JoinRoomResponse {
	return JoinRoomResponse{
		&block.JoinRoomResponse{
			Code:     code,
			Position: position,
		},
	}
}
