package storage

import (
	"github.com/pinfake/pes6go/data/block"
)

type Account struct {
	Id      uint32
	Key     string
	Hash    []byte
	Players [3]uint32
}

type Storage interface {
	CreateAccount(account *Account) (uint32, error)
	Login(account *Account) (*Account, error)
	CreatePlayer(account *Account, position byte, player *block.Player) (uint32, error)
	GetServerNews() []*block.News
	GetRankUrls() []*block.RankUrl
	GetAccountPlayers(account *Account) ([3]*block.Player, error)
	GetGroupInfo(id uint32) *block.GroupInfo
	GetPlayerSettings(id uint32) *block.PlayerSettings
	GetPlayer(id uint32) (*block.Player, error)
	GetLobbies(serverId uint32) []*block.Lobby
}
