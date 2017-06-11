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

type RankUrlsBody struct {
	rankUrls []RankUrlInternal
}

func (info RankUrl) buildInternal() RankUrlInternal {
	var internal RankUrlInternal
	internal.rtype = byte(info.Rtype)
	internal.unknown = 0
	copy(internal.url[:], info.Url)

	return internal
}

func (info RankUrlsBody) GetBytes() []byte {
	buf := bytes.Buffer{}
	for _, rankUrl := range info.rankUrls {
		binary.Write(&buf, binary.BigEndian, rankUrl)
	}
	return buf.Bytes()
}

func (info RankUrls) GetBlock(query uint16) Block {
	body := RankUrlsBody{}
	for _, rankUrl := range info.RankUrls {
		body.rankUrls = append(body.rankUrls, rankUrl.buildInternal())
	}
	return NewBlock(
		query, body,
	)
}
