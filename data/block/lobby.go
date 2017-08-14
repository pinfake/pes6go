package block

import (
	"github.com/pinfake/pes6go/data/types"
)

type Lobbies struct {
	Lobbies []*Lobby
}

type Lobby struct {
	Type       byte
	Name       string
	NumClients uint16
	Rooms      *types.IdMap
}

func GetRoomsSlice(rooms *types.IdMap) []*Room {
	rooms.RLock()
	defer rooms.RUnlock()
	arr := make([]*Room, len(rooms.Data))

	i := 0
	for _, value := range rooms.Data {
		arr[i] = value.(*Room)
		i++
	}

	return arr
}

func (l Lobby) RemoveRoom(roomId uint32) {
	l.Rooms.Delete(roomId)
}

type LobbyInternal struct {
	Type       byte
	Name       [32]byte
	NumClients uint16
}

type LobbiesInternal struct {
	NumLobbies      uint16
	LobbiesInternal []LobbyInternal
}

func (info Lobbies) buildInternal() PieceInternal {
	internals := LobbiesInternal{
		NumLobbies: uint16(len(info.Lobbies)),
	}

	for _, server := range info.Lobbies {
		var internal LobbyInternal
		internal.Type = server.Type
		copy(internal.Name[:], server.Name)
		internal.NumClients = server.NumClients
		internals.LobbiesInternal = append(internals.LobbiesInternal, internal)
	}

	return internals
}
