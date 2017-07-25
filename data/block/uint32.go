package block

import "encoding/binary"

type Uint32 struct {
	Value uint32
}

type Uint32Internal struct {
	Value uint32
}

func (info Uint32) buildInternal() PieceInternal {
	var internal Uint32Internal
	internal.Value = info.Value
	return internal
}

func NewUint32(b *Block) Uint32 {
	return Uint32{
		Value: binary.BigEndian.Uint32(b.Body.GetBytes()),
	}
}
