package blocks

type Void struct {
	Body
}

func (void Void) GetBytes() []byte {
	ret := [0]byte{}
	return ret[:]
}
