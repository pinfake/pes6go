package storage

import "github.com/pinfake/pes6go/data/block"

type Storage interface {
	GetServerNews() []block.News
	GetRankUrls() []block.RankUrl
}
