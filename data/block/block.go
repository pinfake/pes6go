package block

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"crypto/md5"

	"github.com/pinfake/pes6go/network"
)

const dtLayout = "2006-01-15 15:04:05"
const headerSize = 24

type Header struct {
	Query    uint16
	Size     uint16
	Sequence uint32
	Hash     [16]byte // Sixserver says this is the md5 of data
}

type Body interface {
	GetBytes() []byte
}

type Block struct {
	Header Header
	Body   Body
}

type GenericBody struct {
	Data []byte
}

func (body GenericBody) GetBytes() []byte {
	return body.Data
}

func NewHeader(query uint16, size uint16, hash [16]byte) Header {
	return Header{Query: query, Size: size, Hash: hash}
}

func NewBlock(query uint16, body Body) Block {
	return Block{NewHeader(
		query,
		uint16(len(body.GetBytes())),
		md5.Sum(body.GetBytes()),
	), body}
}

func (b Block) GetBytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, b.Header)
	buf.Write(b.Body.GetBytes())
	return buf.Bytes()
}

func ReadBlock(data []byte) (*Block, error) {
	if len(data) < headerSize {
		return nil, errors.New("No Header found")
	}
	decoded := network.Mutate(data)
	var buf = bytes.NewBuffer(decoded[0:headerSize])
	header := Header{}
	err := binary.Read(buf, binary.BigEndian, &header)
	if err != nil {
		panic(err)
	}

	if len(decoded) < int(headerSize+header.Size) {
		fmt.Printf("%d headersize % x", header.Size, decoded)
		return nil, fmt.Errorf(
			"Smaller body than header said, header: %d, body: %d",
			header.Size, len(decoded),
		)
	}
	return &Block{
		header,
		GenericBody{decoded[headerSize : headerSize+header.Size]},
	}, nil
}
