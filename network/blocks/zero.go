package blocks;

import "encoding/binary"

type Zero struct {
    Body
}

func (zero Zero) getData() []byte {
    // TODO: reserve space for the byte array?
    ret := []byte{};
    binary.LittleEndian.PutUint32(ret, uint32(0))
    return ret;
}