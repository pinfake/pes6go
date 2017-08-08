package server

import (
	"log"
	"os"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/storage"
)

type AccountingServer struct {
	config  ServerConfig
	storage storage.Storage
}

var accountingHandlers = map[uint16]Handler{
	0x3010: Profiles,
	0x3020: CreateProfile,
	0x3040: AccountingPlayerInfo,
	0x3050: GroupInfo,
	0x3060: QueryPlayerId,
	0x3070: Unknown3070,
	0x308a: PlayerSettings,
	0x3090: Unknown3090,
	0x3100: Unknown3100,
	0x3120: Unknown3120,
}

func NewAccountingServerHandler(stor storage.Storage) AccountingServer {
	return AccountingServer{
		storage: stor,
		config:  ServerConfig{},
	}
}

func (s AccountingServer) Handlers() map[uint16]Handler {
	return accountingHandlers
}

func (s AccountingServer) Storage() storage.Storage {
	return s.storage
}

func (s AccountingServer) Config() ServerConfig {
	return s.config
}

func (s AccountingServer) Data() interface{} {
	return nil
}

func CreateProfile(s *Server, b *block.Block, c *Connection) message.Message {
	playerCreate := block.NewPlayerCreate(b)
	player := block.NewPlayer(playerCreate.Name)
	responseCode := block.Ok
	_, err := s.Storage().CreatePlayer(c.Account, playerCreate.Position, player)
	if err != nil {
		responseCode = block.ServiceUnavailableError
	}

	return message.PlayerCreateResponse{
		uint32(responseCode),
	}
}

func PlayerSettings(s *Server, b *block.Block, _ *Connection) message.Message {
	playerId := block.NewUint32(b)
	return message.NewPlayerSettingsMessage(
		playerId.Value, s.Storage().GetPlayerSettings(playerId.Value),
	)
}

func Unknown3120(_ *Server, _ *block.Block, _ *Connection) message.Message {
	return message.NewUnknown3120Message()
}

func Unknown3100(_ *Server, _ *block.Block, _ *Connection) message.Message {
	return message.NewUnknown3100Message()
}

func Unknown3070(_ *Server, _ *block.Block, _ *Connection) message.Message {
	return message.NewUnknown3070Message()
}

func Unknown3090(_ *Server, _ *block.Block, _ *Connection) message.Message {
	return message.NewUnknown3090Message()
}

func GroupInfo(s *Server, b *block.Block, _ *Connection) message.Message {
	groupId := block.NewUint32(b)
	return message.NewGroupInfoMessage(
		s.Storage().GetGroupInfo(groupId.Value),
	)
}

func AccountingPlayerInfo(s *Server, b *block.Block, c *Connection) message.Message {
	playerId := block.NewUint32(b)

	player, err := s.Storage().GetPlayer(playerId.Value)
	if err != nil {
		s.Log(c, "Unable to get player %d: %s", playerId.Value, err)
		return nil
	}
	return message.NewAccountingPlayerInfoMessage(
		block.PlayerExtended{player},
	)
}

func QueryPlayerId(_ *Server, b *block.Block, _ *Connection) message.Message {
	_ = block.NewUint32(b)
	return message.NewPlayerIdResponseMessage()
}

func Profiles(s *Server, _ *block.Block, c *Connection) message.Message {
	players, err := s.Storage().GetAccountPlayers(c.Account)
	if err != nil {
		s.Log(c, "Unable to players for account %d: %s", c.Account.Id, err)
		return nil
	}
	return message.NewAccountProfilesMessage(
		block.AccountPlayers{
			players,
		},
	)
}

func StartAccounting(stor storage.Storage) {
	s := NewServer(
		log.New(os.Stdout, "Accounting: ", log.LstdFlags),
		NewAccountingServerHandler(stor),
	)
	s.Serve(12881)
}
