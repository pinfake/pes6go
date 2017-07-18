package server

import (
	"fmt"

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

func Init(_ Server, _ block.Block, _ *Connection) message.Message {
	return message.AccountingInit{}
}

func KeepAlive(_ Server, _ block.Block, _ *Connection) message.Message {
	return message.KeepAlive{}
}

func Disconnect(_ Server, _ block.Block, _ *Connection) message.Message {
	return nil
}

func Login(s Server, b block.Block, c *Connection) message.Message {
	auth := block.NewAthentication(b)
	fmt.Println("I am handling login")
	fmt.Printf("key: % x\n", auth.Key)
	fmt.Printf("password: % x\n", auth.Password)
	fmt.Printf("unknown: % x\n", auth.Unknown)
	fmt.Printf("roster: % x\n", auth.RosterHash)

	bl, _ := blowfish.NewCipher(BlowfishKey)
	decrypter := ecb.NewECBDecrypter(bl)
	dst := make([]byte, len(auth.Key))
	decrypter.CryptBlocks(dst, auth.Key)

	fmt.Printf("cd key decoded: %s\n", dst)

	c.AccountId = s.GetStorage().FindAccount(
		string(dst), auth.Password,
	)

	return message.LoginResponse{
		block.Ok,
	}
}

func SelectPlayer(s Server, b block.Block, c *Connection) message.Message {
	playerSelected := block.NewPlayerSelected(b)
	player := s.GetStorage().GetAccountProfiles(c.AccountId)[playerSelected.Position]
	return message.NewPlayerExtraSettingsMessage(
		block.PlayerExtraSettings{
			PlayerId: player.Id,
		},
	)
}

func ServerLobbies(s Server, _ block.Block, _ *Connection) message.Message {
	a, _ := strconv.ParseUint(s.GetConfig()["serverId"], 10, 32)
	return message.NewLobbiesMessage(
		block.Lobbies{
			s.GetStorage().GetLobbies(
				uint32(a),
			),
		},
	)
}

func JoinLobby(_ Server, b block.Block, _ *Connection) message.Message {
	playerIp := block.NewJoinLobby(b)
	fmt.Printf("%+v\n", playerIp)
	return message.IpInfoResponse{}
}
