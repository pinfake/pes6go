package message

import "github.com/pinfake/pes6go/data/block"

type LeaveRoomResponse struct {
}

func (LeaveRoomResponse) GetBlocks() []*block.Block {
	return block.GetBlocks(0x432b, block.Uint32{0})
}

func NewLeaveRoomResponse() LeaveRoomResponse {
	return LeaveRoomResponse{}
}
