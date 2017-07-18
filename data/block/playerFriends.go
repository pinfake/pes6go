package block

type PlayerFriends struct {
}

type PlayerFriendsInternal struct {
}

func (info PlayerFriends) buildInternal() PieceInternal {
	var internal PlayerFriendsInternal
	return internal
}
