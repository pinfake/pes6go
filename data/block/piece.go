package block

import (
	"bytes"
	"encoding/binary"
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

func getBlocksFromInternals(query uint16, internals []PieceInternal) []*Block {
	var blocks []*Block

	data := getBytesFromInternals(internals)

	for len(data) > MAX_BLOCK_DATA_SIZE {
		chunk := data[:MAX_BLOCK_DATA_SIZE]
		data = data[MAX_BLOCK_DATA_SIZE:]
		blocks = append(blocks, NewBlock(query, GenericBody{chunk}))
	}

	blocks = append(blocks, NewBlock(query, GenericBody{data}))

	return blocks
}

func GetBytes(b PieceInternal) []byte {
	buf := new(bytes.Buffer)

	// I'll use reflection here just in case there is slices
	v := reflect.ValueOf(b)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Slice {
			for j := 0; j < field.Len(); j++ {
				binary.Write(buf, binary.BigEndian, field.Index(j).Interface())
			}
		} else {
			binary.Write(buf, binary.BigEndian, field.Interface())
		}
	}

	return buf.Bytes()
}

func GetBlocks(query uint16, pieces []Piece) []*Block {
	var internals []PieceInternal
	for _, piece := range pieces {
		internals = append(internals, piece.buildInternal())
	}
	return getBlocksFromInternals(query, internals)
}

func GetPieces(data reflect.Value) []Piece {
	if data.Kind() != reflect.Slice {
		return []Piece{data.Interface().(Piece)}
	}
	ret := make([]Piece, data.Len())
	for i := 0; i < data.Len(); i++ {
		ret[i] = data.Index(i).Interface().(Piece)
	}
	return ret
}
