package server

import (
	"fmt"

	"log"
	"os"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/storage"
)

type AccountingServer struct {
	logger      *log.Logger
	connections Connections
}

var accountingHandlers = map[uint16]Handler{
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

func NewAccountingServer() AccountingServer {
	return AccountingServer{
		logger:      log.New(os.Stdout, "Accounting: ", log.LstdFlags),
		connections: NewConnections(),
	}
}

func (s AccountingServer) Handlers() map[uint16]Handler {
	return accountingHandlers
}

func (s AccountingServer) Storage() storage.Storage {
	return storage.Forged{}
}

func (s AccountingServer) Config() ServerConfig {
	return ServerConfig{}
}

func (s AccountingServer) Connections() Connections {
	return s.connections
}

func (s AccountingServer) Logger() *log.Logger {
	return s.logger
}

func CreateProfile(s Server, b block.Block, _ *Connection) message.Message {
	playerCreate := block.NewPlayerCreate(b)
	s.Storage().CreatePlayer(
		playerCreate.Position,
		playerCreate.Name,
	)

	return message.PlayerCreateResponse{
		block.Ok,
	}
}

func PlayerSettings(s Server, b block.Block, _ *Connection) message.Message {
	playerId := block.NewUint32(b)
	return message.NewPlayerSettingsMessage(
		playerId.Value, s.Storage().GetPlayerSettings(playerId.Value),
	)
}

func Unknown3120(_ Server, _ block.Block, _ *Connection) message.Message {
	return message.NewUnknown3120Message()
}

func Unknown3100(_ Server, _ block.Block, _ *Connection) message.Message {
	return message.NewUnknown3100Message()
}

func Unknown3070(_ Server, _ block.Block, _ *Connection) message.Message {
	return message.NewUnknown3070Message()
}

func Unknown3090(_ Server, _ block.Block, _ *Connection) message.Message {
	return message.NewUnknown3090Message()
}

func GroupInfo(s Server, b block.Block, _ *Connection) message.Message {
	groupId := block.NewUint32(b)
	return message.NewGroupInfoMessage(
		s.Storage().GetGroupInfo(groupId.Value),
	)
}

func PlayerGroupInfo(s Server, b block.Block, _ *Connection) message.Message {
	playerId := block.NewUint32(b)
	return message.NewPlayerGroupMessage(
		s.Storage().GetPlayerGroup(playerId.Value),
	)
}

func QueryPlayerId(_ Server, b block.Block, _ *Connection) message.Message {
	_ = block.NewUint32(b)
	return message.NewPlayerIdResponseMessage()
}

func Profiles(s Server, _ block.Block, _ *Connection) message.Message {
	return message.NewAccountProfilesMessage(
		block.AccountPlayers{
			s.Storage().GetAccountProfiles(0),
		},
	)
}

func StartAccounting() {
	fmt.Println("Accounting Server starting")
	Serve(NewAccountingServer(), 12881)
}