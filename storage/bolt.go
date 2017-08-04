package storage

import (
	"bytes"
	"time"

	"encoding/json"

	"fmt"

	"encoding/binary"

	"github.com/boltdb/bolt"
	"github.com/pinfake/pes6go/data/block"
)

type Bolt struct {
	db *bolt.DB
}

func uint32ToBytes(data uint32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, data)
	return buf.Bytes()
}

func NewBolt() (*Bolt, error) {
	db, err := bolt.Open("./pes6godb.bolt", 0600, nil)
	if err != nil {
		return nil, err
	}
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("accounts"))
		tx.CreateBucketIfNotExists([]byte("players"))
		return nil
	})

	return &Bolt{
		db: db,
	}, nil
}

func (b Bolt) CreateAccount(account *Account) (uint32, error) {
	err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("accounts"))
		v := bucket.Get([]byte(account.Key))
		if v != nil {
			return fmt.Errorf("Account %s exists", account.Key)
		}
		id, _ := bucket.NextSequence()
		account.Id = uint32(id)
		buf, err := json.Marshal(account)
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(account.Key), buf)
		return err
	})
	return account.Id, err
}

func (b Bolt) CreatePlayer(account *Account, position byte, player *block.Player) (uint32, error) {
	err := b.db.Update(func(tx *bolt.Tx) error {
		accBckt := tx.Bucket([]byte("accounts"))
		plyBckt := tx.Bucket([]byte("players"))
		id, _ := plyBckt.NextSequence()
		player.Id = uint32(id)
		buf, err := json.Marshal(player)
		if err != nil {
			return err
		}
		err = plyBckt.Put(uint32ToBytes(player.Id), buf)
		if err != nil {
			return err
		} else {
			account.Players[position] = player.Id
			buf, err := json.Marshal(account)
			if err != nil {
				return err
			}
			err = accBckt.Put([]byte(account.Key), buf)
			return err
		}
	})
	return player.Id, err
}

func (b Bolt) GetLobbies(serverId uint32) []block.Lobby {
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

func (b Bolt) GetPlayer(id uint32) *block.Player {
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

func (b Bolt) GetServerNews() []block.News {
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

func (b Bolt) GetRankUrls() []block.RankUrl {
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

func (b Bolt) GetAccountProfiles(id uint32) [3]block.AccountPlayer {
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

func (b Bolt) GetGroupInfo(id uint32) block.GroupInfo {
	return block.GroupInfo{
		Time: time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC),
	}
}

func (b Bolt) GetPlayerSettings(id uint32) block.PlayerSettings {
	return block.PlayerSettings{
		Settings: DefaultPlayerSettings,
	}
}

func (b Bolt) Login(account *Account) (*Account, error) {
	var ret *Account
	fmt.Printf("Here we go, account.Key: %v\n", account.Key)
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("accounts"))
		v := bucket.Get([]byte(account.Key))
		var acc Account
		err := json.Unmarshal(v, &acc)
		if err != nil {
			return err
		}
		if !bytes.Equal(acc.Hash, account.Hash) {
			return fmt.Errorf("Invalid password (hashes don't match)")
		}
		ret = &acc
		return nil
	})
	return ret, err
}
