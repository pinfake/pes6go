package block

type PlayerSelected struct {
	Position byte
	Lang     byte
	Unknown  byte
}

func NewPlayerSelected(b *Block) PlayerSelected {
	return PlayerSelected{
		Position: b.Body.GetBytes()[0],
		Lang:     b.Body.GetBytes()[1],
		Unknown:  b.Body.GetBytes()[2],
	}
}
