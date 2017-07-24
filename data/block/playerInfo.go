package block

type PlayerInfo struct {
	*Player
}

type PlayerInfoInternal struct {
	Zero           uint32
	Id             uint32
	Name           [48]byte
	GroupId        uint32
	GroupName      [48]byte
	GroupStatus    byte
	Division       byte
	Points         uint32
	Category       uint16
	MatchesPlayed  uint16
	Victories      uint16
	Defeats        uint16
	Draws          uint16
	WinningStreak  uint16
	BestStreak     uint16
	Disconnections uint16
	GoalsScored    uint32
	GoalsReceived  uint32
	Comment        [256]byte
	Position       uint32
	// uint16's could have to do with medals
	Unknown2  [12]byte
	Lang      uint16
	LastTeams [5]uint16
}

func (info PlayerInfo) buildInternal() PieceInternal {
	var internal PlayerInfoInternal
	internal.Id = info.Id
	copy(internal.Name[:], info.Name)
	internal.GroupId = info.GroupId
	copy(internal.GroupName[:], info.GroupName)
	internal.GroupStatus = info.GroupStatus
	internal.Division = info.Division
	internal.Points = info.Points
	internal.Category = info.Category
	internal.MatchesPlayed = info.MatchesPlayed
	internal.Victories = info.Victories
	internal.Defeats = info.Defeats
	internal.Draws = info.Draws
	internal.WinningStreak = info.WinningStreak
	internal.BestStreak = info.BestStreak
	internal.Disconnections = info.Disconnections
	internal.GoalsScored = info.GoalsScored
	internal.GoalsReceived = info.GoalsReceived
	copy(internal.Comment[:], info.Comment)
	internal.Position = info.Position
	internal.Lang = info.Lang
	internal.LastTeams = info.Teams
	return internal
}
