package blocks

import "encoding/binary"

type Zero struct {
	Body
}

func (zero Zero) getData() []byte {
	// TODO: reserve space for the byte array?
	ret := [4]byte{}
	binary.BigEndian.PutUint32(ret[:], uint32(0))
	return ret[:]
}
