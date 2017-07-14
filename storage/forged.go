package storage

import (
	"time"

	"github.com/pinfake/pes6go/data/block"
)

type Forged struct {
}

func (_ Forged) GetServerNews() []block.News {
	return []block.News{
		{
			Time:  time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC),
			Title: "Mariano Powered:",
			Text: "Es el vecino el que elige al alcalde y es el alcalde el que quiere " +
				"que sean los vecinos el alcalde",
		},
		{
			Time:  time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC),
			Title: "Mariano Powered:",
			Text: "Es el vecino el que elige al alcalde y es el alcalde el que quiere " +
				"que sean los vecinos el alcalde",
		},
	}
}

func (_ Forged) GetRankUrls() []block.RankUrl {
	return []block.RankUrl{
		{0, "http://pes6web.winning-eleven.net/pes6e2/ranking/we10getrank.html"},
		{1, "https://pes6web.winning-eleven.net/pes6e2/ranking/we10getgrprank.html"},
		{2, "http://pes6web.winning-eleven.net/pes6e2/ranking/we10RankingWeek.html"},
		{3, "https://pes6web.winning-eleven.net/pes6e2/ranking/we10GrpRankingWeek.html"},
		{4, "https://pes6web.winning-eleven.net/pes6e2/ranking/we10RankingCup.html"},
		{5, "http://www.pes6j.net/server/we10getgrpboard.html"},
		{6, "http://www.pes6j.net/server/we10getgrpinvitelist.html"},
	}
}
