package block

import "fmt"

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

type RoomShort struct {
	Room
}

type RoomShortInternal struct {
	Players [4]RoomPlayerShortInternal
}

type RoomPlayerShortInternal struct {
	Id uint32
	Position byte
	Participation byte
}

func (info RoomShort) buildInternal() PieceInternal {
	for i := 0; i < 4; i++ {
		if len(info.Players) > i {

		} else {

		}
	}
	return nil
}

type RoomPlayer struct {
	Id            uint32
	Team          byte
	Spectator     byte
	Participation byte
}

type RoomTeam struct {
	Id          uint16
	GoalsByPart [5]byte
}

type RoomPlayerLink struct {
	Player        *Player
	Position      byte
	Participation byte
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
	internal.Color = info.Participation
	return internal
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
				Id:            0,
				Team:          0xff,
				Spectator:     0,
				Participation: 0xff,
			}
		}
		internal.Players[i] = RoomPlayerInternal{
			Id:        player.Id,
			Owner:     owner,
			Team:      player.Team,
			Spectator: player.Spectator,
			Position:  byte(i),
			Color:     player.Participation,
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
	if info.Players[i].Participation == 0xff {
		// Set participation
		info.Players[i].Participation = info.getNextAvailableParticipation()
	} else {
		info.Players[i].Participation = 0xff
		info.moveDownParticipations(i)
	}
	return info.Players[i].Participation, nil
}

func (info Room) moveDownParticipations(i int) {
	for _, player := range info.Players[i:] {
		if player.Participation != 0xff {
			player.Participation--
		}
	}
}

// TODO: Has to be synchronized, i probably missed many more
func (info Room) getNextAvailableParticipation() byte {
	var participation byte = 0
	for _, player := range info.Players {
		if player.Participation != 0xff {
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
		Id:            player.Id,
		Team:          0xff,
		Spectator:     0,
		Participation: 0xff,
	}
}
