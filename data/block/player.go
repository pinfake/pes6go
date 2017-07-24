package block

type Player struct {
	Id       uint32
	Category uint16
	Points   uint32

	MatchPoints   uint32
	MatchesPlayed uint16
	Victories     uint16
	Defeats       uint16
	Draws         uint16

	WinningStreak  uint16
	BestStreak     uint16
	Disconnections uint16
	Division       byte

	Teams         [5]uint16
	GoalsScored   uint32
	GoalsReceived uint32
	TimePlayed    uint32
	LastLogin     uint32
	Position      uint32
	OldCategory   uint32
	OldPoints     uint32

	Name     string
	Comment  string
	Lang     uint16
	Settings []byte
	LoggedIn bool
	Admin    int

	RoomId uint32

	GroupId     uint32
	GroupName   string
	GroupStatus byte
}

type PlayerInternal struct {
	Id          uint32
	Name        [48]byte
	GroupId     uint32
	GroupName   [48]byte
	GroupStatus byte
	Division    byte
	RoomId      uint32
	Unknown1    uint16
	// My self from the past says this could tell about the group level
	Unknown2 uint16
	Category uint16
	// My self from the past says this could tell about the group level as well
	Unknown3  uint16
	Victories uint16
	Defeats   uint16
	Draws     uint16
	Unknown4  [3]byte
}

func (info Player) buildInternal() PieceInternal {
	var internal PlayerInternal
	internal.Id = info.Id
	copy(internal.Name[:], info.Name)
	internal.GroupId = info.GroupId
	copy(internal.GroupName[:], info.GroupName)
	internal.GroupStatus = info.GroupStatus
	internal.Division = info.Division
	internal.RoomId = info.RoomId
	internal.Unknown1 = 0x0000
	internal.Unknown2 = 0x3fff
	internal.Category = info.Category
	internal.Unknown3 = 0x3fff
	internal.Victories = info.Victories
	internal.Defeats = info.Defeats
	internal.Draws = info.Draws

	return internal
}
