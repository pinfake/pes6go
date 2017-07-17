package storage

import "github.com/pinfake/pes6go/data/block"

type Storage interface {
	CreateAccount(key string, password string) uint32
	FindAccount(key string, hash []byte) uint32
	CreatePlayer(position byte, name string)
	GetServerNews() []block.News
	GetRankUrls() []block.RankUrl
	GetAccountProfiles(id uint32) [3]block.AccountPlayer
	GetPlayerGroup(id uint32) block.PlayerGroup
	GetGroupInfo(id uint32) block.GroupInfo
	GetPlayerSettings(id uint32) block.PlayerSettings
}
