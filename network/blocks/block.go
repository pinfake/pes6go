package blocks;

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

func NewHeader(query uint16, size uint16) Header {
    return Header{Query:query, Size:size};
}

func NewBlock(query uint16, body Body) Block {
    return Block{NewHeader(query, uint16(len(body.getData()))), body};
}