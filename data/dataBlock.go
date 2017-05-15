package data

import "github.com/pinfake/pes6go/network/blocks"

type dataBlock interface {
	getBlock(queryId uint16) blocks.Block
}
