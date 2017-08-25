package server

import (
	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/storage"
	"golang.org/x/crypto/blowfish"
)

var handlers = map[uint16]Handler{
	0x0003: Disconnect,
	0x0005: KeepAlive,
	0x3001: Init,
	0x3003: Login,
	0x3080: PlayerFriends,
	0x4100: SelectPlayer,
	0x4200: ServerLobbies,
	0x4202: JoinLobby,
}

func Init(_ *Server, _ *block.Block, _ *Connection) message.Message {
	return message.NewAccountingInit()
}

func KeepAlive(_ *Server, _ *block.Block, _ *Connection) message.Message {
	return message.KeepAlive{}
}

func LeaveLobby(s *Server, _ *block.Block, c *Connection) {
	sendToLobby(s.connections, c.LobbyId, message.NewLeaveLobby(c.Player.Id))
	c.LobbyId = 0xff
}

func Disconnect(s *Server, _ *block.Block, c *Connection) message.Message {
	if c.Player != nil {
		LeaveRoom(s, nil, c)
		LeaveLobby(s, nil, c)
	}

	return nil
}

func Login(s *Server, b *block.Block, c *Connection) message.Message {
	auth := block.NewAthentication(b)

	bl, _ := blowfish.NewCipher(BlowfishKey)
	decrypter := ecb.NewECBDecrypter(bl)
	dst := make([]byte, len(auth.Key))
	decrypter.CryptBlocks(dst, auth.Key)

	s.Log(c, "LOGIN -> Key: %s, Pass: %x, Roster: %x", dst, auth.Password, auth.RosterHash)
	acc := storage.Account{
		Key:  string(dst[:20]),
		Hash: auth.Password,
	}
	found, err := s.Storage().Login(&acc)
	code := block.Ok
	if err != nil {
		s.Log(c, "Cannot login: %s", err)
		code = block.ServiceUnavailableError
	} else {
		c.Account = found
	}

	return message.NewLoginResponse(uint32(code))
}

func SelectPlayer(s *Server, b *block.Block, c *Connection) message.Message {
	playerSelected := block.NewPlayerSelected(b)
	players, err := s.Storage().GetAccountPlayers(c.Account)
	if err != nil {
		s.Log(c, "Unable to get player profiles for %s: %s", c.Account.Id, err)
		return nil
	}
	c.Player = players[playerSelected.Position]
	c.Player.ResetRoomData()
	return message.NewPlayerExtraSettingsMessage(
		block.PlayerExtraSettings{
			PlayerId: c.Player.Id,
		},
	)
}

func ServerLobbies(s *Server, _ *block.Block, _ *Connection) message.Message {
	return message.NewLobbies(
		&block.Lobbies{s.lobbies},
	)
}

func JoinLobby(s *Server, b *block.Block, c *Connection) message.Message {
	joinLobby := block.NewJoinLobby(b)
	c.LobbyId = joinLobby.LobbyId
	c.Player.Link = &block.Link{
		Ip1:   joinLobby.Ip1,
		Port1: joinLobby.Port1,
		Ip2:   joinLobby.Ip2,
		Port2: joinLobby.Port2,
	}
	sendToLobby(s.connections, c.LobbyId, message.NewPlayerJoinedLobbyMessage(*c.Player))
	return message.NewJoinLobbyResponse(block.Ok)
}

func PlayerFriends(_ *Server, _ *block.Block, _ *Connection) message.Message {
	return message.NewPlayerFriendsMessage(block.PlayerFriends{})
}
