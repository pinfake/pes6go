package data

import (
	"bytes"
	"encoding/binary"

	"github.com/pinfake/pes6go/network/blocks"
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

type RankUrlBlock struct {
	rankUrls []RankUrlInternal
}

func (info RankUrl) buildInternal() RankUrlInternal {
	var internal RankUrlInternal
	internal.rtype = byte(info.Rtype)
	internal.unknown = 0
	copy(internal.url[:], info.Url)

	return internal
}

func (info RankUrlBlock) GetBytes() []byte {
	buf := bytes.Buffer{}
	for _, rankUrl := range info.rankUrls {
		binary.Write(&buf, binary.BigEndian, rankUrl)
	}
	return buf.Bytes()
}

func (info RankUrls) GetBlock(query uint16) blocks.Block {
	block := RankUrlBlock{}
	for _, rankUrl := range info.RankUrls {
		block.rankUrls = append(block.rankUrls, rankUrl.buildInternal())
	}
	return blocks.NewBlock(
		query, block,
	)
}
