package network;

import (
    "errors"
    "bytes"
    "encoding/binary"
)

const headerSize = 24;

var key = []byte{0xa6, 0x77, 0x95, 0x7c}

func Mutate(data []byte) ([] byte) {
    decoded := []byte{};
    i := 0;
    j := 0;
    for i < len(data) {
        decoded = append(decoded, data[i]^key[j]);
        j++;
        i++;
        if j%4 == 0 {
            j = 0;
        }
    }
    return decoded;
}

func Read(data []byte) (Block, error) {
    if len(data) < headerSize {
        return Block{}, errors.New("No header found");
    }
    decoded := Mutate(data[0:headerSize]);
    buf := bytes.NewBuffer(decoded);
    header := Header{};
    err := binary.Read(buf, binary.LittleEndian, &header);
    if err != nil {
        panic(err);
    }
    return Block{header, []byte{}}, nil;
}