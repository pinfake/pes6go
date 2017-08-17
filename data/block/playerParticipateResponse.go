package block

type PlayerParticipateResponse struct {
	Code          uint32
	Selection     byte
	Participation byte
}

type PlayerParticipateResponseInternal struct {
	Code          uint32
	Selection     byte
	Participation byte
}

func (info PlayerParticipateResponse) buildInternal() PieceInternal {
	var internal PlayerParticipateResponseInternal
	internal.Code = info.Code
	internal.Selection = info.Selection
	internal.Participation = info.Participation
	return internal
}
