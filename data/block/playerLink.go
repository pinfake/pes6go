package block

type PlayerLink struct {
	Player *Player
}

type PlayerLinkInternal struct {
	Unknown1 uint32
	Ip1      [16]byte
	Port1    uint16
	Ip2      [16]byte
	Port2    uint16
	PlayerId uint32
}

func (info PlayerLink) buildInternal() PieceInternal {
	var internal PlayerLinkInternal
	internal.Unknown1 = 0
	copy(internal.Ip1[:], []byte(info.Player.Link.Ip1))
	copy(internal.Ip2[:], []byte(info.Player.Link.Ip2))
	internal.Port1 = info.Player.Link.Port1
	internal.Port2 = info.Player.Link.Port2
	internal.PlayerId = info.Player.Id
	return internal
}
