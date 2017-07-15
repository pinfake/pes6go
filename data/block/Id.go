package block

import "encoding/binary"

type Id struct {
	Id uint32
}

func NewId(b Block) Id {
	return Id{
		Id: binary.BigEndian.Uint32(b.body.GetBytes()),
	}
}
