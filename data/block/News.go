package block

import (
	"time"
)

type News struct {
	Time  time.Time
	Title string
	Text  string
}

type NewsInternal struct {
	header [6]byte
	time   [19]byte
	title  [64]byte
	text   [935]byte
}

func (info News) buildInternal() PieceInternal {
	var internal NewsInternal
	copy(internal.header[:], []byte{0x00, 0x00, 0x03, 0x10, 0x01, 0x00})
	copy(internal.time[:], info.Time.Format(dtLayout))
	copy(internal.title[:], info.Title)
	copy(internal.text[:], info.Text)

	return internal
}
