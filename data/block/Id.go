package block

import "encoding/binary"

type Id struct {
	Id uint32
}

type IdInternal struct {
	id uint32
}

func (info Id) buildInternal() PieceInternal {
	var internal IdInternal
	internal.id = info.Id
	return internal
}

func NewId(b Block) Id {
	return Id{
		Id: binary.BigEndian.Uint32(b.body.GetBytes()),
	}
}
