package blocks

import (
	"bytes"
	"encoding/binary"
	"time"
)

type ServerTime struct {
	Time time.Time
}

type timeInternal struct {
	time uint32
}

func (m ServerTime) buildInternal() timeInternal {
	internal := timeInternal{
		time: uint32(m.Time.Unix() / 1000),
	}
	return internal
}

func (m ServerTime) GetBytes() []byte {
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, m.buildInternal())
	return buf.Bytes()
}
