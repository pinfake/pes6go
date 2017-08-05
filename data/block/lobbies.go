package block

type Lobbies struct {
	Lobbies []*Lobby
}

type Lobby struct {
	Type       byte
	Name       string
	NumClients uint16
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
