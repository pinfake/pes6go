package block

type AccountPlayers struct {
	Players [3]AccountPlayer
}

type AccountPlayer struct {
	Position      byte
	Id            uint32
	Name          string
	TimePlayed    uint32
	Division      byte
	Points        uint32
	Category      uint16
	MatchesPlayed uint16
}

type AccountPlayersInternal struct {
	zero                   uint32
	AccountPlayersInternal [3]AccountPlayerInternal
}

type AccountPlayerInternal struct {
	position      byte
	id            uint32
	name          [32]byte
	unknown       [16]byte
	timePlayed    uint32
	division      byte
	points        uint32
	category      uint16
	matchesPlayed uint16
}

func (info AccountPlayers) buildInternal() PieceInternal {
	internals := AccountPlayersInternal{
		zero: 0,
	}
	for i, player := range info.Players {
		var internal AccountPlayerInternal
		internal.position = byte(i)
		internal.id = player.Id
		copy(internal.name[:], player.Name)
		internal.unknown = [16]byte{
			0x00, 0x60, 0xdc, 0x2e, 0x34, 0xd7, 0x69, 0x64,
			0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}
		internal.timePlayed = player.TimePlayed
		internal.division = player.Division
		internal.points = player.Points
		internal.category = player.Category
		internal.matchesPlayed = player.MatchesPlayed
		internals.AccountPlayersInternal[i] = internal
	}
	return internals
}
