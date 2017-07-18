package block

type Server struct {
	Stype      int
	Name       string
	Ip         string
	Port       int
	NumClients int
}

type ServerInternal struct {
	Unknown1   [7]byte
	Stype      byte
	Name       [32]byte
	Ip         [15]byte
	Port       uint16
	NumClients uint16
	Unknown2   [2]byte
}

func (info Server) buildInternal() PieceInternal {
	var internal ServerInternal
	copy(internal.Unknown1[:], []byte{
		0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00,
	})
	internal.Stype = byte(info.Stype)
	copy(internal.Name[:], info.Name)
	copy(internal.Ip[:], info.Ip)
	internal.Port = uint16(info.Port)
	internal.NumClients = uint16(info.NumClients)
	copy(internal.Unknown2[:], []byte{0x00, 0x00})

	return internal
}
