package block

type PlayerGroup struct {
	PlayerName string
	GroupId    uint32
	GroupName  string
}

type PlayerGroupInternal struct {
	Zero       uint32
	PlayerName [48]byte
	GroupId    uint32
	GroupName  [48]byte
	Unknown    [294]byte
}

func (info PlayerGroup) buildInternal() PieceInternal {
	var internal PlayerGroupInternal
	internal.Zero = 0
	copy(internal.PlayerName[:], info.PlayerName)
	internal.GroupId = info.GroupId
	copy(internal.GroupName[:], info.GroupName)

	return internal
}
