package block

import (
	"time"
)

type ServerTime struct {
	Time time.Time
}

type ServerTimeInternal struct {
	time uint32
}

func (info ServerTime) buildInternal() PieceInternal {
	internal := ServerTimeInternal{
		time: uint32(info.Time.Unix()),
	}
	return internal
}
