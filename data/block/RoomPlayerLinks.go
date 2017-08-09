package block

type RoomPlayerLinks struct {
	PlayerLinks []*RoomPlayerLink
}

type RoomPlayerLink struct {
	Player   *Player
	Position byte
	Color    byte
}

type RoomPlayerLinkInternal struct {
	Unknown1  [32]byte
	Ip1       [16]byte
	Port1     uint16
	Ip2       [16]byte
	Port2     uint16
	PlayerId  uint32
	Position1 byte // Why two positions? idk
	Position2 byte
	Color     byte
}

func (info RoomPlayerLink) buildInternal() PieceInternal {
	var internal RoomPlayerLinkInternal
	internal.Unknown1 = [32]byte{
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x80,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xc0,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
	}
	copy(internal.Ip1[:], []byte(info.Player.Link.Ip1))
	copy(internal.Ip2[:], []byte(info.Player.Link.Ip2))
	internal.Port1 = info.Player.Link.Port1
	internal.Port2 = info.Player.Link.Port2
	internal.PlayerId = info.Player.Id
	internal.Position1 = info.Position
	internal.Position2 = info.Position
	internal.Color = info.Color
	return internal
}
