package storage

import (
	"bytes"
	"time"

	"encoding/json"

	"fmt"

	"strconv"

	"os"

	"github.com/boltdb/bolt"
	"github.com/pinfake/pes6go/data/block"
)

type Bolt struct {
	db *bolt.DB
}

func uint32ToBytes(data uint32) []byte {
	return []byte(strconv.Itoa(int(data)))
}

func NewBolt() (*Bolt, error) {
	_ = os.Mkdir("./db", 0700)
	db, err := bolt.Open("./db/pes6godb.bolt", 0600, nil)
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

func (b Bolt) GetLobbies(serverId uint32) []*block.Lobby {
	switch serverId {
	case 1:
		return []*block.Lobby{
			{Type: 63, Name: "Lobby 1 Kenobi"},
			{Type: 63, Name: "Lobby 2 testá3"},
			{Type: 63, Name: "Lobby 3 testñ3"},
		}
	case 2:
		return []*block.Lobby{
			{Type: 0x1f},
		}
	default:
		return nil
	}
}

func (b Bolt) GetPlayer(id uint32) (*block.Player, error) {
	var player *block.Player
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("players"))
		v := bucket.Get(uint32ToBytes(id))
		if v == nil {
			return fmt.Errorf("Player with id %d not found", id)
		}
		player = &block.Player{}
		err := json.Unmarshal(v, player)
		if err != nil {
			return err
		}
		return nil
	})

	return player, err
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

func (b Bolt) GetAccountPlayers(account *Account) ([3]*block.Player, error) {
	var players [3]*block.Player
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("players"))
		for i := range account.Players {
			v := bucket.Get(uint32ToBytes(account.Players[i]))
			var player block.Player
			if v == nil {
				player = block.Player{}
			} else {
				err := json.Unmarshal(v, &player)
				if err != nil {
					return err
				}
			}
			players[i] = &player
		}
		return nil
	})

	return players, err
}

func (b Bolt) GetGroupInfo(id uint32) *block.GroupInfo {
	return &block.GroupInfo{
		Time: time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC),
	}
}

func (b Bolt) GetPlayerSettings(id uint32) *block.PlayerSettings {
	return &block.PlayerSettings{
		Settings: DefaultPlayerSettings,
	}
}

func (b Bolt) Login(account *Account) (*Account, error) {
	var ret *Account
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("accounts"))
		v := bucket.Get([]byte(account.Key))
		if v == nil {
			return fmt.Errorf("Account key '%s' not found!", account.Key)
		}
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
