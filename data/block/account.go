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
	Zero                   uint32
	AccountPlayersInternal [3]AccountPlayerInternal
}

type AccountPlayerInternal struct {
	Position      byte
	Id            uint32
	Name          [32]byte
	Unknown       [16]byte
	TimePlayed    uint32
	Division      byte
	Points        uint32
	Category      uint16
	MatchesPlayed uint16
}

func (info AccountPlayers) buildInternal() PieceInternal {
	internals := AccountPlayersInternal{
		Zero: 0,
	}
	for i, player := range info.Players {
		var internal AccountPlayerInternal
		internal.Position = byte(i)
		internal.Id = player.Id
		copy(internal.Name[:], player.Name)
		internal.Unknown = [16]byte{
			0x00, 0x60, 0xdc, 0x2e, 0x34, 0xd7, 0x69, 0x64,
			0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}
		internal.TimePlayed = player.TimePlayed
		internal.Division = player.Division
		internal.Points = player.Points
		internal.Category = player.Category
		internal.MatchesPlayed = player.MatchesPlayed
		internals.AccountPlayersInternal[i] = internal
	}
	return internals
}
