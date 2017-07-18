package block

import (
	"encoding/binary"
)

type JoinLobby struct {
	LobbyId byte
	Ip1     string
	Port1   uint16
	Ip2     string
	Port2   uint16
}

type JoinLobbyInternal struct {
	LobbyId byte
	Ip1     [16]byte
	Port1   uint16
	Ip2     [16]byte
	Port2   uint16
	unknown [2]byte
}

func (info JoinLobby) buildInternal() PieceInternal {
	var internal JoinLobbyInternal
	internal.LobbyId = info.LobbyId
	copy(internal.Ip1[:], []byte(info.Ip1))
	internal.Port1 = info.Port1
	copy(internal.Ip2[:], []byte(info.Ip2))
	internal.Port2 = info.Port2
	return internal
}

func NewJoinLobby(b Block) JoinLobby {
	bytes := b.Body.GetBytes()
	return JoinLobby{
		LobbyId: bytes[0],
		Ip1:     string(bytes[1:17]),
		Port1:   binary.BigEndian.Uint16(bytes[17:19]),
		Ip2:     string(bytes[19:35]),
		Port2:   binary.BigEndian.Uint16(bytes[35:37]),
	}
}
