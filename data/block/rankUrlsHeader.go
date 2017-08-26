package block

type RankUrlsHeader struct {
}

type RankUrlsInternal struct {
	Unknown [8]byte
}

func (info RankUrlsHeader) buildInternal() PieceInternal {
	return RankUrlsInternal{
		Unknown: [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55},
	}
}
