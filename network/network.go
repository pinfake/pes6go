package network;

import (
    "errors"
    "bytes"
    "encoding/binary"
)

const headerSize = 24;

var key = []byte {0xa6, 0x77, 0x95, 0x7c}

type Header struct {
    query uint16;
    size uint16;
    sequence uint32;
    unknown [16]byte;
}

type Message struct {
	header Header;
	body []byte;
}

func Mutate(data []byte) ([] byte) {
	decoded := []byte {};
	i := 0;
    j := 0;
    for i < len(data) {
        decoded = append(decoded, data[i]^key[j]);
        j++;
        i++;
        if j % 4 == 0 {
            j = 0;
        }
    }
	return decoded;
}

func Read(data []byte) (Message, error) {
	if len(data) < headerSize {
		return Message{}, errors.New("No header found");
	}
	decoded := Mutate(data[0:headerSize - 1]);
    buf := bytes.NewBuffer(decoded);
    header := Header{};
    err := binary.Read(buf, binary.BigEndian, &header);
    if err != nil {
        panic(err);
    }
    return Message{header, []byte{}}, nil;
}