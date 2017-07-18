package block

type PlayerIdResponse struct {
}

type PlayerIdResponseInternal struct {
	unknown1 [8]byte
	code     uint16
	unknown2 [4]byte
}

func (info PlayerIdResponse) buildInternal() PieceInternal {
	internal := PlayerIdResponseInternal{}
	internal.code = 0x0001
	return internal
}
