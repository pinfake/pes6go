package info

type Player struct {
	Id       uint32
	Category uint32
	Points   uint32

	MatchPoints   uint32
	MatchesPlayed uint32
	Victories     uint32
	Defeats       uint32
	Draws         uint32

	WinningStreak  uint32
	BestStreak     uint32
	Disconnections uint32
	Division       uint32

	Teams         []uint32
	GoalsScored   uint32
	GoalsReceived uint32
	TimePlayed    uint32
	LastLogin     uint32
	Position      uint32
	OldCategory   uint32
	OldPoints     uint32

	Name     string
	Comment  string
	Settings []byte
	LoggedIn bool
	Admin    int

	GroupId uint32
}
