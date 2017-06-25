package block

type Authentication struct {
	Key        []byte
	Password   []byte
	Unknown    []byte
	RosterHash []byte
}

func NewAthentication(b Block) Authentication {
	return Authentication{
		Key:        b.body.GetBytes()[:32],
		Password:   b.body.GetBytes()[32:48],
		Unknown:    b.body.GetBytes()[48:58],
		RosterHash: b.body.GetBytes()[58:74],
	}
}
