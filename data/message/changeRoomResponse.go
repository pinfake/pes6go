package message

import "github.com/pinfake/pes6go/data/block"

type ChangeRoomResponse struct {
	Code uint32
}

func (m ChangeRoomResponse) GetBlocks() []*block.Block {
	return block.GetBlocks(0x434e, []block.Piece{
		block.Uint32{m.Code},
	})
}

func NewChangeRoomResponse(code uint32) ChangeRoomResponse {
	return ChangeRoomResponse{
		Code: code,
	}
}
