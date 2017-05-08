package messages

import "github.com/pinfake/pes6go/network/blocks"

type Message interface {
	getBlocks() []blocks.Block
}
