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

var blowfishKey = []byte{
	0x27, 0x50, 0x1f, 0xd0, 0x4e, 0x6b, 0x82, 0xc8,
	0x31, 0x02, 0x4d, 0xac, 0x5c, 0x63, 0x05, 0x22,
	0x19, 0x74, 0xde, 0xb9, 0x38, 0x8a, 0x21, 0x90,
	0x1d, 0x57, 0x6c, 0xbb, 0xe2, 0xf3, 0x77, 0xef,
	0x23, 0xd7, 0x54, 0x86, 0x01, 0x0f, 0x37, 0x81,
	0x9a, 0xfe, 0x6c, 0x32, 0x1a, 0x01, 0x46, 0xd2,
	0x15, 0x44, 0xec, 0x36, 0x5b, 0xf7, 0x28, 0x9a,
}

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

func QueryPlayerId(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.NewPlayerIdResponseMessage(block.PlayerIdOk)
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

	bl, _ := blowfish.NewCipher(blowfishKey)
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
