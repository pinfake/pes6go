package block

import (
	"fmt"
)

type Room struct {
	Id          uint32
	Type        byte
	Phase       byte
	Name        string
	Time        byte
	Players     []*Player
	Teams       [2]RoomTeam
	HasPassword byte
	Password    string
	MatchType   byte
	ChatLevel   byte
}

type RoomParticipation Room

type RoomParticipationInternal struct {
	Players [4]RoomPlayerParticipationInternal
}

type RoomPlayerParticipationInternal struct {
	Id            uint32
	Position      byte
	Participation byte
}

func (info RoomParticipation) buildInternal() PieceInternal {
	var internal RoomParticipationInternal
	for i := 0; i < 4; i++ {
		if len(info.Players) > i {
			internal.Players[i] = RoomPlayerParticipationInternal{
				Id:            info.Players[i].Id,
				Position:      byte(i),
				Participation: info.Players[i].RoomData.Participation,
			}
		} else {
			internal.Players[i] = RoomPlayerParticipationInternal{
				Id:            0,
				Position:      byte(i),
				Participation: 0xff,
			}
		}
	}
	return internal
}

type RoomTeam struct {
	Id          uint16
	GoalsByPart [5]byte
}

type RoomPlayerLinks Room

type RoomPlayerLinksInternal struct {
	Internals []RoomPlayerLinkInternal
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

func (info RoomPlayerLinks) buildInternal() PieceInternal {
	var thisInternal RoomPlayerLinksInternal
	thisInternal.Internals = make([]RoomPlayerLinkInternal, len(info.Players))
	for i := 0; i < len(info.Players); i++ {
		var internal = &thisInternal.Internals[i]
		internal.Unknown1 = [32]byte{
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x80,
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xc0,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		}
		copy(internal.Ip1[:], []byte(info.Players[i].Link.Ip1))
		copy(internal.Ip2[:], []byte(info.Players[i].Link.Ip2))
		internal.Port1 = info.Players[i].Link.Port1
		internal.Port2 = info.Players[i].Link.Port2
		internal.PlayerId = info.Players[i].Id
		internal.Position1 = byte(i)
		internal.Position2 = byte(i)
		internal.Color = info.Players[i].RoomData.Participation
	}
	return thisInternal
}

type RoomPlayerInternal struct {
	Id            uint32
	Owner         byte
	Unknown       byte
	Team          byte
	Spectator     byte
	Position      byte
	Participation byte
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
		if len(info.Players) > i {
			internal.Players[i] = RoomPlayerInternal{
				Id:            info.Players[i].Id,
				Owner:         owner,
				Team:          info.Players[i].RoomData.Team,
				Spectator:     info.Players[i].RoomData.Spectator,
				Position:      byte(i),
				Participation: info.Players[i].RoomData.Participation,
			}
			fmt.Printf("id %x, position: %x, participation, %x\n",
				info.Players[i].Id,
				byte(i),
				info.Players[i].RoomData.Participation)
		} else {
			internal.Players[i] = RoomPlayerInternal{
				Id:            0,
				Owner:         0,
				Team:          0xff,
				Spectator:     0,
				Position:      byte(i),
				Participation: 0xff,
			}
		}

	}

	internal.RoomTeams = info.Teams
	internal.HasPassword = info.HasPassword
	internal.MatchType = info.MatchType
	internal.ChatLevel = info.ChatLevel
	return internal
}

func (info *Room) HasPlayers() bool {
	fmt.Printf("Pero tengo a gente??? %v\n", info.Players)
	return info.GetNumPlayers() > 0
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
	if info.Players[i].RoomData.Participation == 0xff {
		// Set participation
		info.Players[i].RoomData.Participation = info.getNextAvailableParticipation()
	} else {
		info.Players[i].RoomData.Participation = 0xff
		info.moveDownParticipations(i)
	}
	return info.Players[i].RoomData.Participation, nil
}

func (info Room) moveDownParticipations(i int) {
	for _, player := range info.Players[i:] {
		if player.RoomData.Participation != 0xff {
			player.RoomData.Participation--
		}
	}
}

// TODO: Has to be synchronized, i probably missed many more
func (info Room) getNextAvailableParticipation() byte {
	var participation byte = 0
	for _, player := range info.Players {
		if player.RoomData.Participation != 0xff {
			participation++
		}
	}
	return participation
}

func (info *Room) RemovePlayer(playerId uint32) error {
	i, err := info.getPlayerIdx(playerId)
	if err != nil {
		// Log something here, the player wasnt found
		return err
	}
	info.Players = append(info.Players[:i], info.Players[i+1:]...)
	fmt.Printf("Borro al tío y me queda %v\n", info.Players)
	return nil
}

func (info *Room) AddPlayer(player *Player) byte {
	info.Players = append(info.Players, player)
	return byte(info.GetNumPlayers() - 1)
}

func (info *Room) GetNumPlayers() int {
	return len(info.Players)
}
