package block

type PlayerSelected struct {
	Position byte
}

func NewPlayerSelected(b Block) PlayerSelected {
	return PlayerSelected{
		Position: b.Body.GetBytes()[0],
	}
}
