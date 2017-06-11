package block

import (
	"bytes"
	"encoding/binary"
	"time"
)

type ServerTime struct {
	Time time.Time
}

type ServerTimeBody struct {
	time uint32
}

func (info ServerTime) buildInternal() ServerTimeBody {
	internal := ServerTimeBody{
		time: uint32(info.Time.Unix()),
	}
	return internal
}

func (body ServerTimeBody) GetBytes() []byte {
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, body)
	return buf.Bytes()
}

func (info ServerTime) GetBlock(query uint16) Block {
	return NewBlock(query, info.buildInternal())
}
