package block

import "fmt"

type RoomPlayer struct {
	Id        uint32
	Team      byte
	Spectator byte
	Color     byte
}

type RoomTeam struct {
	Id          uint16
	GoalsByPart [5]byte
}

type Room struct {
	Id          uint32
	Type        byte
	Phase       byte
	Name        string
	Time        byte
	Players     []*RoomPlayer
	Teams       [2]RoomTeam
	HasPassword byte
	Password    string
	MatchType   byte
	ChatLevel   byte
}

type RoomPlayerInternal struct {
	Id        uint32
	Owner     byte
	Unknown   byte
	Team      byte
	Spectator byte
	Position  byte
	Color     byte
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
	for i := 0; i < 4; i++ {
		var owner byte
		if i == 0 {
			owner = 0x01
		}
		var player *RoomPlayer
		if len(info.Players) > i {
			player = info.Players[i]
		} else {
			player = &RoomPlayer{
				Id:        0,
				Team:      0xff,
				Spectator: 0,
				Color:     0xff,
			}
		}
		internal.Players[i] = RoomPlayerInternal{
			Id:        player.Id,
			Owner:     owner,
			Team:      player.Team,
			Spectator: player.Spectator,
			Position:  byte(i),
			Color:     player.Color,
		}
	}

	internal.RoomTeams = info.Teams
	internal.HasPassword = info.HasPassword
	internal.MatchType = info.MatchType
	internal.ChatLevel = info.ChatLevel
	return internal
}

func (info Room) HasPlayers() bool {
	return len(info.Players) > 0
}

func (info Room) getPlayerIdx(playerId uint32) (int, error) {
	for i, player := range info.Players {
		if player.Id == playerId {
			return i, nil
		}
	}
	return 0, fmt.Errorf("player not found")
}

func (info Room) ToggleParticipation(playerId uint32) (byte, error) {
	i, err := info.getPlayerIdx(playerId)
	if err != nil {
		// Log something here, the player wasnt found
		return 0, err
	}
	if info.Players[i].Color == 0xff {
		// Set participation
		info.Players[i].Color = info.getNextAvailableParticipation()
	} else {
		info.Players[i].Color = 0xff
		info.moveDownParticipations(i)
	}
	return info.Players[i].Color, nil
}

func (info Room) moveDownParticipations(i int) {
	for _, player := range info.Players[i:] {
		if player.Color != 0xff {
			player.Color--
		}
	}
}

// TODO: Has to be synchronized, i probably missed many more
func (info Room) getNextAvailableParticipation() byte {
	var participation byte = 0
	for _, player := range info.Players {
		if player.Color != 0xff {
			participation++
		}
	}
	return participation
}

func (info Room) RemovePlayer(playerId uint32) error {
	i, err := info.getPlayerIdx(playerId)
	if err != nil {
		// Log something here, the player wasnt found
		return err
	}
	info.Players = append(info.Players[:i], info.Players[i+1:]...)
	return nil
}

func NewRoomPlayer(player *Player) *RoomPlayer {
	return &RoomPlayer{
		Id:        player.Id,
		Team:      0xff,
		Spectator: 0,
		Color:     0xff,
	}
}
