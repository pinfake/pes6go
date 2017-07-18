package block

type PlayerIdResponse struct {
}

type PlayerIdResponseInternal struct {
	Unknown1 [8]byte
	Code     uint16
	Unknown2 [4]byte
}

func (info PlayerIdResponse) buildInternal() PieceInternal {
	internal := PlayerIdResponseInternal{}
	internal.Code = 0x0001
	return internal
}
