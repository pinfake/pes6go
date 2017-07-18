package menu

import (
	"fmt"

	"golang.org/x/crypto/blowfish"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/server"
	"github.com/pinfake/pes6go/storage"
)

type MenuServer struct {
	storage storage.Storage
}

var handlers = map[uint16]server.Handler{
	0x0003: Disconnect,
	0x0005: KeepAlive,
	0x3001: Init,
	0x3003: Login,
	0x3080: PlayerFriends,
	0x4100: SelectPlayer,
	0x4200: MenuServers,
	0x4202: IpInfo,
}

func PlayerFriends(_ server.Server, b block.Block, _ *server.Connection) message.Message {
	return message.NewPlayerFriendsMessage(block.PlayerFriends{})
}

func IpInfo(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.IpInfoResponse{}
}

func MenuServers(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.NewMenuServersMessage(
		block.MenuServers{
			[1]block.MenuServer{
				{8, "MENU03-SP/", 0},
			},
		},
	)
}

func SelectPlayer(s server.Server, b block.Block, c *server.Connection) message.Message {
	playerSelected := block.NewPlayerSelected(b)
	player := s.(MenuServer).storage.GetAccountProfiles(c.AccountId)[playerSelected.Position]
	return message.NewPlayerExtraSettingsMessage(
		block.PlayerExtraSettings{
			PlayerId: player.Id,
		},
	)
}

func Init(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.AccountingInit{}
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

	c.AccountId = s.(MenuServer).storage.FindAccount(
		string(dst), auth.Password,
	)

	return message.LoginResponse{
		block.Ok,
	}
}

func (s MenuServer) GetHandlers() map[uint16]server.Handler {
	return handlers
}

func KeepAlive(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return message.KeepAlive{}
}

func Disconnect(_ server.Server, _ block.Block, _ *server.Connection) message.Message {
	return nil
}

func Start() {
	fmt.Println("Menu Server starting")
	server.Serve(MenuServer{
		storage: storage.Forged{},
	}, 12882)
}
