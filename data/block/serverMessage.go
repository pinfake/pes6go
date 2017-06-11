package block

import (
	"bytes"
	"encoding/binary"
	"time"
)

const dtLayout = "2006-01-15 15:04:05"

type ServerMessage struct {
	Time  time.Time
	Title string
	Text  string
	dataBlock
}

type ServerMessageBody struct {
	header [6]byte
	time   [19]byte
	title  [64]byte
	text   [128]byte
}

func (info ServerMessage) buildInternal() ServerMessageBody {
	var internal ServerMessageBody
	copy(internal.header[:], []byte{0x00, 0x00, 0x03, 0x10, 0x01, 0x00})
	copy(internal.time[:], info.Time.Format(dtLayout))
	copy(internal.title[:], info.Title)
	copy(internal.text[:], info.Text)

	return internal
}

func (info ServerMessageBody) GetBytes() []byte {
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, info)
	return buf.Bytes()
}

func (info ServerMessage) GetBlock(query uint16) Block {
	return NewBlock(
		query, info.buildInternal(),
	)
}
