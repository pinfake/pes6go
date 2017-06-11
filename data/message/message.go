package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type Message interface {
	GetBlocks() []block.Block
}
