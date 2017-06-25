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

func (info ServerInternal) getBytes() []byte {
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, info)
	return buf.Bytes()
}

func (info Servers) GetBlocks(query uint16) []Block {
	bits := []BlockBit{}
	for _, server := range info.Servers {
		bits = append(bits, server.buildInternal())
	}
	return GetBlocks(query, bits)
}
