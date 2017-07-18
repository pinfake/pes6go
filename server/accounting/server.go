package accounting

import (
	"fmt"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/server"
	"github.com/pinfake/pes6go/storage"
	"golang.org/x/crypto/blowfish"
)

type AccountingServer struct {
	storage storage.Storage
}

var handlers = map[uint16]server.Handler{
	0x0003: Disconnect,
	0x0005: KeepAlive,
	0x3001: Init,
	0x3003: Login,
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

func CreateProfile(s server.Server, b block.Block, _ *server.Connection) message.Message {
	playerCreate := block.NewPlayerCreate(b)
	s.(AccountingServer).storage.CreatePlayer(
		playerCreate.Position,
		playerCreate.Name,
	)

	return message.PlayerCreateResponse{
		block.Ok,
	}
}

func PlayerSettings(s server.Server, b block.Block, _ *server.Connection) message.Message {
	playerId := block.NewId(b)
	return message.NewPlayerSettingsMessage(
		playerId.Id, s.(AccountingServer).storage.GetPlayerSettings(playerId.Id),
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
	groupId := block.NewId(b)
	return message.NewGroupInfoMessage(
		s.(AccountingServer).storage.GetGroupInfo(groupId.Id),
	)
}

func PlayerGroupInfo(s server.Server, b block.Block, _ *server.Connection) message.Message {
	playerId := block.NewId(b)
	return message.NewPlayerGroupMessage(
		s.(AccountingServer).storage.GetPlayerGroup(playerId.Id),
	)
}

func QueryPlayerId(_ server.Server, b block.Block, _ *server.Connection) message.Message {
	_ = block.NewId(b)
	return message.NewPlayerIdResponseMessage()
}

func Profiles(s server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.NewAccountProfilesMessage(
		block.AccountPlayers{
			s.(AccountingServer).storage.GetAccountProfiles(0),
		},
	)
}

func Login(s server.Server, b block.Block, c *server.Connection) message.Message {
	auth := block.NewAthentication(b)
	fmt.Println("I am handling login")
	fmt.Printf("key: % x\n", auth.Key)
	fmt.Printf("password: % x\n", auth.Password)
	fmt.Printf("unknown: % x\n", auth.Unknown)
	fmt.Printf("roster: % x\n", auth.RosterHash)

	bl, _ := blowfish.NewCipher(server.BlowfishKey)
	decrypter := ecb.NewECBDecrypter(bl)
	dst := make([]byte, len(auth.Key))
	decrypter.CryptBlocks(dst, auth.Key)

	fmt.Printf("cd key decoded: %s\n", dst)

	c.AccountId = s.(AccountingServer).storage.FindAccount(
		string(dst), auth.Password,
	)

	return message.LoginResponse{
		block.Ok,
	}
}

func Init(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.AccountingInit{}
}

func KeepAlive(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.KeepAlive{}
}

func Disconnect(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return nil
}

func Start() {
	fmt.Println("Accounting Server starting")
	server.Serve(AccountingServer{
		storage.Forged{},
	}, 12881)
}
