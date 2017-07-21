package server

import (
	"strconv"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"golang.org/x/crypto/blowfish"
)

var handlers = map[uint16]Handler{
	0x0003: Disconnect,
	0x0005: KeepAlive,
	0x3001: Init,
	0x3003: Login,
	0x4100: SelectPlayer,
	0x4200: ServerLobbies,
	0x4202: JoinLobby,
}

func Init(_ *Server, _ block.Block, _ *Connection) message.Message {
	return message.AccountingInit{}
}

func KeepAlive(_ *Server, _ block.Block, _ *Connection) message.Message {
	return message.KeepAlive{}
}

func Disconnect(_ *Server, _ block.Block, _ *Connection) message.Message {
	return nil
}

func Login(s *Server, b block.Block, c *Connection) message.Message {
	auth := block.NewAthentication(b)

	bl, _ := blowfish.NewCipher(BlowfishKey)
	decrypter := ecb.NewECBDecrypter(bl)
	dst := make([]byte, len(auth.Key))
	decrypter.CryptBlocks(dst, auth.Key)

	s.Log(c, "LOGIN -> Key: %s, Pass: %x, Roster: %x", dst, auth.Password, auth.RosterHash)

	c.AccountId = s.Storage().FindAccount(
		string(dst[:20]), auth.Password,
	)

	code := block.Ok
	if c.AccountId < 1 {
		code = block.ServiceUnavailableError
	}
	return message.LoginResponse{
		uint32(code),
	}
}

func SelectPlayer(s *Server, b block.Block, c *Connection) message.Message {
	playerSelected := block.NewPlayerSelected(b)
	playerProfile := s.Storage().GetAccountProfiles(c.AccountId)[playerSelected.Position]
	player := s.Storage().GetPlayer(playerProfile.Id)
	c.Player = &player
	return message.NewPlayerExtraSettingsMessage(
		block.PlayerExtraSettings{
			PlayerId: playerProfile.Id,
		},
	)
}

func ServerLobbies(s *Server, _ block.Block, _ *Connection) message.Message {
	a, _ := strconv.ParseUint(s.Config()["serverId"], 10, 32)

	return message.NewLobbiesMessage(
		block.Lobbies{
			s.Storage().GetLobbies(
				uint32(a),
			),
		},
	)
}

func JoinLobby(s *Server, b block.Block, c *Connection) message.Message {
	playerIp := block.NewJoinLobby(b)
	s.Log(c, "JOIN LOBBY -> %+v", playerIp)
	return message.IpInfoResponse{}
}
