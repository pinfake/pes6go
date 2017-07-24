package block

type RoomPlayer struct {
	Id        uint32
	Team      byte
	Spectator byte
	Ready     byte
}

type RoomTeam struct {
	Id          uint16
	GoalsByPart [5]byte
}

type Room struct {
	Id        uint32
	Type      byte
	Phase     byte
	Name      string
	Time      byte
	Players   [4]*RoomPlayer
	Teams     [2]RoomTeam
	Password  string
	MatchType byte
	ChatLevel byte
}

type RoomPlayerInternal struct {
	Id        uint32
	Owner     byte
	Unknown   byte
	Team      byte
	Spectator byte
	Position  byte
	Ready     byte
}

type RoomInternal struct {
	Id          uint32
	Type        byte
	Phase       byte
	Name        [64]byte
	Time        byte
	Players     [4]RoomPlayerInternal
	RoomTeams   [2]RoomTeam
	Unknown1    byte
	HasPassword byte
	MatchType   byte
	ChatLevel   byte
	Unknown2    byte
	Unknown3    byte
}

func (info Room) buildInternal() PieceInternal {
	var internal RoomInternal
	internal.Id = info.Id
	internal.Type = info.Type
	internal.Phase = info.Phase
	copy(internal.Name[:], info.Name)
	internal.Time = info.Time
	for i, player := range info.Players {
		var owner byte
		if i == 0 {
			owner = 1
		}
		internal.Players[i] = RoomPlayerInternal{
			Id:        player.Id,
			Owner:     owner,
			Team:      player.Team,
			Spectator: player.Spectator,
			Position:  byte(i),
			Ready:     player.Ready,
		}
	}

	internal.RoomTeams = info.Teams
	if info.Password != "" {
		internal.HasPassword = 1
	}
	internal.MatchType = info.MatchType
	internal.ChatLevel = info.ChatLevel
	return internal
}
