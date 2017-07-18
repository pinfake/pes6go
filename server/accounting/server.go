package accounting

import (
	"fmt"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/server"
	"github.com/pinfake/pes6go/storage"
)

type AccountingServer struct {
}

var handlers = map[uint16]server.Handler{
	0x3010: Profiles,
	0x3020: CreateProfile,
	0x3040: PlayerGroupInfo,
	0x3050: GroupInfo,
	0x3060: QueryPlayerId,
	0x3070: Unknown3070,
	0x308a: PlayerSettings,
	0x3090: Unknown3090,
	0x3100: Unknown3100,
	0x3120: Unknown3120,
}

func (s AccountingServer) GetHandlers() map[uint16]server.Handler {
	return handlers
}

func (s AccountingServer) GetStorage() storage.Storage {
	return storage.Forged{}
}

func (s AccountingServer) GetConfig() server.ServerConfig {
	return server.ServerConfig{}
}

func CreateProfile(s server.Server, b block.Block, _ *server.Connection) message.Message {
	playerCreate := block.NewPlayerCreate(b)
	s.GetStorage().CreatePlayer(
		playerCreate.Position,
		playerCreate.Name,
	)

	return message.PlayerCreateResponse{
		block.Ok,
	}
}

func PlayerSettings(s server.Server, b block.Block, _ *server.Connection) message.Message {
	playerId := block.NewUint32(b)
	return message.NewPlayerSettingsMessage(
		playerId.Value, s.GetStorage().GetPlayerSettings(playerId.Value),
	)
}

func Unknown3120(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.NewUnknown3120Message()
}

func Unknown3100(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.NewUnknown3100Message()
}

func Unknown3070(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.NewUnknown3070Message()
}

func Unknown3090(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.NewUnknown3090Message()
}

func GroupInfo(s server.Server, b block.Block, _ *server.Connection) message.Message {
	groupId := block.NewUint32(b)
	return message.NewGroupInfoMessage(
		s.GetStorage().GetGroupInfo(groupId.Value),
	)
}

func PlayerGroupInfo(s server.Server, b block.Block, _ *server.Connection) message.Message {
	playerId := block.NewUint32(b)
	return message.NewPlayerGroupMessage(
		s.GetStorage().GetPlayerGroup(playerId.Value),
	)
}

func QueryPlayerId(_ server.Server, b block.Block, _ *server.Connection) message.Message {
	_ = block.NewUint32(b)
	return message.NewPlayerIdResponseMessage()
}

func Profiles(s server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.NewAccountProfilesMessage(
		block.AccountPlayers{
			s.GetStorage().GetAccountProfiles(0),
		},
	)
}

func Start() {
	fmt.Println("Accounting Server starting")
	server.Serve(AccountingServer{}, 12881)
}
