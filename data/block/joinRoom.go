package block

import "encoding/binary"

type JoinRoomRequest struct {
	Id       uint32
	Password string
}

type JoinRoomResponse struct {
	Code     uint32
	Position byte
}

type JoinRoomResponseInternal struct {
	Code     uint32
	Position byte
}

func (info JoinRoomResponse) buildInternal() PieceInternal {
	return JoinRoomResponseInternal{
		Code:     info.Code,
		Position: info.Position,
	}
}

func NewJoinRoomRequest(b *Block) JoinRoomRequest {
	bytes := b.Body.GetBytes()
	return JoinRoomRequest{
		Id:       binary.BigEndian.Uint32(bytes[0:4]),
		Password: string(bytes[4:]),
	}
}
