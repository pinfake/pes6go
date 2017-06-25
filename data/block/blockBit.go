package block

const MAX_BLOCK_DATA_SIZE = 256

type BlockBit interface {
	getBytes() (bytes []byte)
}

func GetBlocks(query uint16, bits []BlockBit) []Block {
	var blocks []Block
	var buffer []byte

	for _, bit := range bits {
		bytes := bit.getBytes()
		if len(bytes)+len(buffer) > MAX_BLOCK_DATA_SIZE {
			blocks = append(
				blocks,
				NewBlock(query, GenericBody{buffer}),
			)
			buffer = nil
		}
		buffer = append(buffer, bytes...)
	}
	blocks = append(
		blocks,
		NewBlock(query, GenericBody{buffer}),
	)
	return blocks
}
