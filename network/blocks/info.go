package blocks

import (
	"bytes"
	"encoding/binary"
	"time"
)

const dtLayout = "2006-01-15 15:04:05"

type Info struct {
	Time  time.Time
	Title string
	Text  string
}

type infoInternal struct {
	header [6]byte
	time   [20]byte
	title  [64]byte
	text   [128]byte
}

func (info Info) buildInternal() infoInternal {
	var internal infoInternal
	copy(internal.header[:], []byte{0x00, 0x00, 0x03, 0x10, 0x01, 0x00})
	copy(internal.time[:], info.Time.Format(dtLayout))
	copy(internal.title[:], info.Title)
	copy(internal.text[:], info.Text)

	return internal
}

func (info Info) GetBytes() []byte {
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, info.buildInternal())
	return buf.Bytes()
}
