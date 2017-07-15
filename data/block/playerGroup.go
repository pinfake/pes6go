package block

type PlayerGroup struct {
	PlayerName string
	GroupId    uint32
	GroupName  string
}

type PlayerGroupInternal struct {
	zero       uint32
	playerName [48]byte
	groupId    uint32
	groupName  [48]byte
	unknown    [294]byte
}

func (info PlayerGroup) buildInternal() PieceInternal {
	var internal PlayerGroupInternal
	internal.zero = 0
	copy(internal.playerName[:], info.PlayerName)
	internal.groupId = info.GroupId
	copy(internal.groupName[:], info.GroupName)

	return internal
}
