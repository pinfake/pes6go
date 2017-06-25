package block

import (
	"bytes"
	"encoding/binary"
)

const MAX_BLOCK_DATA_SIZE = 512

type PieceInternal interface {
}

type Piece interface {
	buildInternal() PieceInternal
}

func getBlocksFromInternals(query uint16, internals []PieceInternal) []Block {
	var blocks []Block
	var buffer []byte

	for _, internal := range internals {
		data := GetBytes(internal)

		if len(data)+len(buffer) > MAX_BLOCK_DATA_SIZE {
			blocks = append(
				blocks,
				NewBlock(query, GenericBody{buffer}),
			)
			buffer = nil
		}
		buffer = append(buffer, data...)
	}
	blocks = append(
		blocks,
		NewBlock(query, GenericBody{buffer}),
	)
	return blocks
}

func GetBytes(b PieceInternal) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, b)
	return buf.Bytes()
}

func GetBlocks(query uint16, pieces []Piece) []Block {
	var internals []PieceInternal
	for _, piece := range pieces {
		internals = append(internals, piece.buildInternal())
	}
	return getBlocksFromInternals(query, internals)
}