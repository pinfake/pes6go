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
	Header [6]byte
	Time   [19]byte
	Title  [64]byte
	Text   [935]byte
}

func (info News) buildInternal() PieceInternal {
	var internal NewsInternal
	copy(internal.Header[:], []byte{0x00, 0x00, 0x03, 0x10, 0x01, 0x00})
	copy(internal.Time[:], info.Time.Format(dtLayout))
	copy(internal.Title[:], info.Title)
	copy(internal.Text[:], info.Text)

	return internal
}
