package block

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

const MAX_BLOCK_DATA_SIZE = 1024

type Pieces struct {
	pieces []interface{}
}

type PieceInternal interface {
}

type Piece interface {
	buildInternal() PieceInternal
}

func getBytesFromInternals(internals []PieceInternal) []byte {
	var data []byte
	for _, internal := range internals {
		data = append(data, GetBytes(internal)...)
	}

	return data
}

func getBlocksFromInternals(query uint16, internals []PieceInternal) []Block {
	var blocks []Block

	data := getBytesFromInternals(internals)

	for len(data) > MAX_BLOCK_DATA_SIZE {
		chunk := data[:MAX_BLOCK_DATA_SIZE]
		data = data[MAX_BLOCK_DATA_SIZE:]
		blocks = append(blocks, NewBlock(query, GenericBody{chunk}))
	}

	if len(data) > 0 {
		blocks = append(blocks, NewBlock(query, GenericBody{data}))
	}
	return blocks
}

func GetBytes(b PieceInternal) []byte {
	buf := new(bytes.Buffer)
	fmt.Printf("pieceinternal: % x\n", b)
	binary.Write(buf, binary.BigEndian, b)
	fmt.Printf("binarywrite:   % x\n", buf.Bytes())
	return buf.Bytes()
}

func GetBlocks(query uint16, pieces []Piece) []Block {
	var internals []PieceInternal
	for _, piece := range pieces {
		internals = append(internals, piece.buildInternal())
	}
	return getBlocksFromInternals(query, internals)
}

func GetPieces(slice reflect.Value) []Piece {
	if slice.Kind() != reflect.Slice {
		return []Piece{slice.Interface().(Piece)}
	}
	ret := make([]Piece, slice.Len())
	for i := 0; i < slice.Len(); i++ {
		ret[i] = slice.Index(i).Interface().(Piece)
	}
	return ret
}
