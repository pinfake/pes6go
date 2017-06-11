package block

type dataBlock interface {
	getBlock(queryId uint16) Block
}
