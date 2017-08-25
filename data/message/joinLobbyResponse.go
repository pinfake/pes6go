package message

import "github.com/pinfake/pes6go/data/block"

type JoinLobbyResponse struct {
	Code uint32
}

func (data JoinLobbyResponse) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4203, block.Uint32{data.Code})
}

func NewJoinLobbyResponse(code uint32) JoinLobbyResponse {
	return JoinLobbyResponse{code}
}
