package storage

import "github.com/pinfake/pes6go/data/block"

type Storage interface {
	FindAccount(key string, hash []byte) uint32
	GetServerNews() []block.News
	GetRankUrls() []block.RankUrl
	GetAccountProfiles(id uint32) [3]block.AccountPlayer
	GetPlayerGroup(id uint32) block.PlayerGroup
	GetGroupInfo(id uint32) block.GroupInfo
	GetPlayerSettings(id uint32) block.PlayerSettings
}
