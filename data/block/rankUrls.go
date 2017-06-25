package block

type RankUrl struct {
	Rtype int
	Url   string
}

type RankUrlInternal struct {
	rtype   byte
	unknown byte
	url     [128]byte
}

func (info RankUrl) buildInternal() PieceInternal {
	var internal RankUrlInternal
	internal.rtype = byte(info.Rtype)
	internal.unknown = 0
	copy(internal.url[:], info.Url)

	return internal
}
