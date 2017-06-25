package block

import (
	"time"
)

const dtLayout = "2006-01-15 15:04:05"

type ServerMessage struct {
	Time  time.Time
	Title string
	Text  string
}

type ServerMessageInternal struct {
	header [6]byte
	time   [19]byte
	title  [64]byte
	text   [128]byte
}

func (info ServerMessage) buildInternal() PieceInternal {
	var internal ServerMessageInternal
	copy(internal.header[:], []byte{0x00, 0x00, 0x03, 0x10, 0x01, 0x00})
	copy(internal.time[:], info.Time.Format(dtLayout))
	copy(internal.title[:], info.Title)
	copy(internal.text[:], info.Text)

	return internal
}
