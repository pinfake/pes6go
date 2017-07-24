package block

type RankUrl struct {
	Rtype int
	Url   string
}

type RankUrlInternal struct {
	Rtype   byte
	Unknown byte
	Url     [128]byte
}

func (info RankUrl) buildInternal() PieceInternal {
	var internal RankUrlInternal
	internal.Rtype = byte(info.Rtype)
	internal.Unknown = 0
	copy(internal.Url[:], info.Url)

	return internal
}
