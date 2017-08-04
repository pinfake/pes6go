package storage

import (
	"time"

	"bytes"

	"fmt"

	"github.com/pinfake/pes6go/data/block"
)

type Forged struct {
}

func (_ Forged) GetServerNews() []block.News {
	return []block.News{
		{
			Time:  time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC),
			Title: "Mariano Speaks!",
			Text: "Es el vecino el que elige al alcalde y es el alcalde el que quiere " +
				"que sean los vecinos el alcalde",
		},
		{
			Time:  time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC),
			Title: "Mariano Keeps Speaking!",
			Text:  "No he dormido nada, no me pregunten demasiado si hacen el favor",
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

func (_ Forged) Login(account *Account) (*Account, error) {
	if account.Key == "RFLJY34DRE993HX3ER94" && bytes.Equal(
		account.Hash, []byte{
			0xac, 0x04, 0x6e, 0x00, 0x7a, 0x40, 0x06, 0x17,
			0x0e, 0x9a, 0xc7, 0x3f, 0x66, 0x53, 0x31, 0x71,
		}) {
		return &Account{
			Id:      1234,
			Key:     account.Key,
			Hash:    account.Hash,
			Players: [3]uint32{12345, 2345},
		}, nil
	} else {
		return nil, fmt.Errorf("Invalid password (hashes don't match)")
	}
}

func (_ Forged) CreateAccount(account *Account) (uint32, error) {
	return 1234, nil
}

func (_ Forged) CreatePlayer(account *Account, position byte, player *block.Player) (uint32, error) {
	return 12982, nil
}

func (_ Forged) GetLobbies(serverId uint32) []block.Lobby {
	switch serverId {
	case 1:
		return []block.Lobby{
			{63, "Lobby 1 Kenobi", 23},
			{63, "Lobby 2 testá3", 43},
			{63, "Lobby 3 testñ3", 42},
		}
	case 2:
		return []block.Lobby{
			{0x1f, "", 0},
		}
	default:
		return nil
	}
}

func (_ Forged) GetPlayer(id uint32) *block.Player {
	return &block.Player{
		Position:      1,
		Id:            12345,
		Name:          "PadreJohn",
		TimePlayed:    1000,
		Division:      2,
		Points:        0,
		Category:      500,
		MatchesPlayed: 20,
	}
}
