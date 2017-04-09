package network;

import "errors"

var key = []byte {0xa6, 0x77, 0x95, 0x7c}

type header struct {

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

func Decode(data []byte) (Message, error) {
	if len(data) < 24 {
		return Message{}, errors.New("No header found");
	}
	header := decodeChunk(data[0:23]);
    return Message{header,nil}, nil;
}