package network;

import "errors"

var key = []byte {0xa6, 0x77, 0x95, 0x7c}

type header struct {

}

type Message struct {
	header [24]byte;
	body []byte;
}

func Decode(data []byte) (Message, error) {
	if len(data) < 24 {
		return Message{}, errors.New("No header found");
	}

    return Message{}, nil;
}