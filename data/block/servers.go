package block

import (
	"bytes"
	"encoding/binary"
)

type Server struct {
	Stype      int
	Name       string
	Ip         string
	Port       int
	NumClients int
}

type Servers struct {
	Servers []Server
}

type ServerInternal struct {
	unknown1   [7]byte
	stype      byte
	name       [32]byte
	ip         [15]byte
	port       uint16
	numClients uint16
	unknown2   [2]byte
}

type ServerBody struct {
	servers []ServerInternal
}

func (info Server) buildInternal() ServerInternal {
	var internal ServerInternal
	copy(internal.unknown1[:], []byte{
		0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00,
	})
	internal.stype = byte(info.Stype)
	copy(internal.name[:], info.Name)
	copy(internal.ip[:], info.Ip)
	internal.port = uint16(info.Port)
	internal.numClients = uint16(info.NumClients)
	copy(internal.unknown2[:], []byte{0x00, 0x00})

	return internal
}

func (body ServerBody) GetBytes() []byte {
	buf := bytes.Buffer{}
	for _, server := range body.servers {
		binary.Write(&buf, binary.BigEndian, server)
	}
	return buf.Bytes()
}

func (info Servers) GetBlock(query uint16) Block {
	body := ServerBody{}
	for _, server := range info.Servers {
		body.servers = append(body.servers, server.buildInternal())
	}
	return NewBlock(
		query, body,
	)
}
