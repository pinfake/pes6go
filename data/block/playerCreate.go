package block

type PlayerCreate struct {
	Position byte
	Name     string
}

type PlayerCreateInternal struct {
	position byte
	name     [30]byte
}

func (info PlayerCreate) buildInternal() PieceInternal {
	var internal PlayerCreateInternal
	internal.position = info.Position
	copy(internal.name[:], info.Name)
	return internal
}

func NewPlayerCreate(b Block) PlayerCreate {
	return PlayerCreate{
		Position: b.Body.GetBytes()[0],
		Name:     string(b.Body.GetBytes()[1:]),
	}
}
