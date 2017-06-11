package block

import "encoding/binary"

type Zero struct {
	Body
}

func (zero Zero) GetBytes() []byte {
	ret := [4]byte{}
	binary.BigEndian.PutUint32(ret[:], uint32(0))
	return ret[:]
}
