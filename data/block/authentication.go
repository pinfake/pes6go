package block

type Authentication struct {
	Key        []byte
	Password   []byte
	Unknown    []byte
	RosterHash []byte
}

func NewAthentication(b Block) Authentication {
	return Authentication{
		Key:        b.Body.GetBytes()[:32],
		Password:   b.Body.GetBytes()[32:48],
		Unknown:    b.Body.GetBytes()[48:58],
		RosterHash: b.Body.GetBytes()[58:74],
	}
}
