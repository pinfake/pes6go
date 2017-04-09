package network;

import (
    "errors"
    "bytes"
    "encoding/binary"
)

var key = []byte {0xa6, 0x77, 0x95, 0x7c}

type Header struct {
    unknown1 int16;
    size int16;
    unknown2 [20]byte;
}

type Message struct {
	header []byte;
	body []byte;
}

func decodeChunk(data []byte) ([] byte) {
	decoded := []byte {};
	i := 0;
	for j := 0; j < 4; j++ {
		decoded = append(decoded, data[i] ^ key[j]);
		i++;
	}
	return decoded;
}

func Decode(data []byte) (Header, error) {
	if len(data) < 24 {
		return Header{}, errors.New("No header found");
	}
	decoded := decodeChunk(data[0:23]);
    buf := bytes.NewBuffer(decoded);
    header := Header{};
    err := binary.Read(buf, binary.BigEndian, &header);
    if err != nil {
        panic(err);
    }
    return header, nil;
}