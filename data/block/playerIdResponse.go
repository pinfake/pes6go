package block

const PlayerIdOk = 0x01

type PlayerIdResponse struct {
	Code uint16
}

type PlayerIdResponseInternal struct {
	unknown1 [8]byte
	code     uint16
	unknown2 [4]byte
}

func (info PlayerIdResponse) buildInternal() PieceInternal {
	internal := PlayerIdResponseInternal{}
	internal.code = info.Code
	return internal
}
