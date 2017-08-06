package block

type CreateRoom struct {
	Name        string
	HasPassword byte
	Password    string
}

type CreateRoomInternal struct {
	Name        [64]byte
	HasPassword byte
	Password    [16]byte
}

func (info CreateRoom) buildInternal() PieceInternal {
	var internal CreateRoomInternal
	copy(internal.Name[:], []byte(info.Name))
	internal.HasPassword = info.HasPassword
	copy(internal.Password[:], []byte(info.Password))
	return internal
}

func NewCreateRoom(b *Block) CreateRoom {
	bytes := b.Body.GetBytes()
	return CreateRoom{
		Name:        string(bytes[0:64]),
		HasPassword: bytes[64],
		Password:    string(bytes[65:]),
	}
}
