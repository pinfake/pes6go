package info

type LobbyList struct {
	lobbies []Lobby
}

type Lobby struct {
	players []Player
}

func (l Lobby) PlayerById(id uint32) Player {
	return Player{}
}

func (l Lobby) PlayerByName(name string) Player {
	return Player{}
}
