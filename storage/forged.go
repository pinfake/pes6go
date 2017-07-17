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
			Time:  time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC),
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

func (_ Forged) GetAccountProfiles(id uint32) [3]block.AccountPlayer {
	return [3]block.AccountPlayer{
		{
			Position:      0,
			Id:            12345,
			Name:          "PadreJohn",
			TimePlayed:    1000,
			Division:      2,
			Points:        0,
			Category:      500,
			MatchesPlayed: 20,
		},
		{
			Position:      1,
			Id:            2345,
			Name:          "Danilo",
			TimePlayed:    500,
			Division:      1,
			Points:        50000,
			Category:      1000,
			MatchesPlayed: 90,
		},
		{
			Position:      2,
			Id:            0,
			Name:          "",
			TimePlayed:    0,
			Division:      2,
			Points:        0,
			Category:      500,
			MatchesPlayed: 0,
		},
	}
}

func (_ Forged) GetPlayerGroup(id uint32) block.PlayerGroup {
	return block.PlayerGroup{
		PlayerName: "PadreJohn",
		GroupId:    1234,
		GroupName:  "Outlaws in law",
	}
}

func (_ Forged) GetGroupInfo(id uint32) block.GroupInfo {
	return block.GroupInfo{
		Time: time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC),
	}
}

func (_ Forged) GetPlayerSettings(id uint32) block.PlayerSettings {
	return block.PlayerSettings{
		Settings: DefaultPlayerSettings,
	}
}

func (_ Forged) FindAccount(key string, hash []byte) uint32 {
	return 1234
}
