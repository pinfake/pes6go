package block

import (
	"time"
)

type ServerTime struct {
	Time time.Time
}

type ServerTimeInternal struct {
	Time uint32
}

func (info ServerTime) buildInternal() PieceInternal {
	internal := ServerTimeInternal{
		Time: uint32(info.Time.Unix()),
	}
	return internal
}
