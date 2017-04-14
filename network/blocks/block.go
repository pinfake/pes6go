package blocks;

import (
    "errors"
    "bytes"
    "encoding/binary"
    "github.com/pinfake/pes6go/network"
)

const headerSize = 24;

type Header struct {
    Query    uint16;
    Size     uint16;
    Unknown1 uint16;
    Sequence uint16;
    Unknown2 [16]byte;
}

type Body interface {
    getData() []byte
}

type Block struct {
    header Header;
    body Body;
}

type GenericBody struct {
    data []byte;
}

func (body GenericBody) getData() []byte {
    return body.data;
}

func NewHeader(query uint16, size uint16) Header {
    return Header{Query:query, Size:size};
}

func NewBlock(query uint16, body Body) Block {
    return Block{NewHeader(query, uint16(len(body.getData()))), body};
}

func ReadBlock(data []byte) (Block, error){
    if len(data) < headerSize {
        return Block{}, errors.New("No header found");
    }
    decoded := network.Mutate(data);
    var buf = bytes.NewBuffer(decoded[0:headerSize]);
    header := Header{};
    err := binary.Read(buf, binary.LittleEndian, &header);
    if err != nil {
        panic(err);
    }
    return Block{header, GenericBody{decoded[headerSize:header.Size]}}, nil
}