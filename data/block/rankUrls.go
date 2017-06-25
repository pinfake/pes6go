package block

import (
	"bytes"
	"encoding/binary"
)

type RankUrl struct {
	Rtype int
	Url   string
}

type RankUrls struct {
	RankUrls []RankUrl
}

type RankUrlInternal struct {
	rtype   byte
	unknown byte
	url     [128]byte
}

func (info RankUrl) buildInternal() RankUrlInternal {
	var internal RankUrlInternal
	internal.rtype = byte(info.Rtype)
	internal.unknown = 0
	copy(internal.url[:], info.Url)

	return internal
}

func (info RankUrlInternal) getBytes() []byte {
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, info)
	return buf.Bytes()
}

func (info RankUrls) GetBlocks(query uint16) []Block {
	bits := []BlockBit{}
	for _, rankUrl := range info.RankUrls {
		bits = append(bits, rankUrl.buildInternal())
	}
	return GetBlocks(query, bits)
}
