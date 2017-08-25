package message

import "github.com/pinfake/pes6go/data/block"

type ChangeRoomResponse struct {
	Code uint32
}

func (data ChangeRoomResponse) GetBlocks() []*block.Block {
	return block.GetBlocks(0x434e, block.Uint32{data.Code})
}

func NewChangeRoomResponse(code uint32) ChangeRoomResponse {
	return ChangeRoomResponse{code}
}
