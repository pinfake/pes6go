package data

import (
	"bytes"
	"encoding/binary"

	"github.com/pinfake/pes6go/network/blocks"
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

type ServerBlock struct {
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

func (info ServerBlock) GetBytes() []byte {
	buf := bytes.Buffer{}
	for _, server := range info.servers {
		binary.Write(&buf, binary.BigEndian, server)
	}
	return buf.Bytes()
}

func (info Servers) GetBlock(query uint16) blocks.Block {
	block := ServerBlock{}
	for _, server := range info.Servers {
		block.servers = append(block.servers, server.buildInternal())
	}
	return blocks.NewBlock(
		query, block,
	)
}
