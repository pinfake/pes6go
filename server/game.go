package server

import (
	"fmt"

	"log"
	"os"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/storage"
)

type GameServerData struct {
	rooms *IdMap
}

type GameServer struct {
	data    GameServerData
	config  ServerConfig
	storage storage.Storage
}

var gameHandlers = map[uint16]Handler{
	0x308c: Unknown308c,
	0x4102: GamePlayerInfo,
	0x4210: PlayersInLobby,
	0x4300: RoomsInLobby,
	0x4310: CreateRoom,
	0x4400: Chat,
}

func NewGameServerHandler(stor storage.Storage) GameServer {
	return GameServer{
		storage: stor,
		data:    GameServerData{NewIdMap()},
		config: ServerConfig{
			"serverId": "1",
			"lobbies": "[" +
				"{\"Type\":63, \"Name\":\"Lobby #1\"}," +
				"{\"Type\":63, \"Name\":\"Lobby #2\"}" +
				"]",
		},
	}
}

func (s GameServer) Storage() storage.Storage {
	return s.storage
}

func (s GameServer) Handlers() map[uint16]Handler {
	return gameHandlers
}

func (s GameServer) Config() ServerConfig {
	return s.config
}

func (s GameServer) Data() interface{} {
	return s.data
}

func CreateRoom(s *Server, b *block.Block, c *Connection) message.Message {
	createRoom := block.NewCreateRoom(b)
	s.Log(c, "Create room: %v", createRoom)
	room := block.Room{
		Id:          s.Data().(GameServerData).rooms.GetNewId(),
		Name:        createRoom.Name,
		HasPassword: createRoom.HasPassword,
		Password:    createRoom.Password,
		Players: [4]*block.RoomPlayer{
			block.NewRoomPlayer(c.Player),
			block.NewRoomPlayer(&block.Player{}),
			block.NewRoomPlayer(&block.Player{}),
			block.NewRoomPlayer(&block.Player{}),
		},
	}
	s.Data().(GameServerData).rooms.Add(room.Id, room)
	//s.lobbies[c.LobbyId].Rooms = append(s.lobbies[c.LobbyId].Rooms, &room)
	c.Player.RoomId = room.Id
	c.writeMessage(message.NewRoomUpdateMessage(room))
	//c.writeMessage(message.NewPlayerUpdateMessage(*c.Player))
	return message.NewPlayerUpdateMessage(*c.Player)
}

func PlayersInLobby(s *Server, _ *block.Block, c *Connection) message.Message {
	return message.NewPlayersInLobbyMessage(
		playersInLobby(s.connections, c.LobbyId),
	)
}

func RoomsInLobby(s *Server, _ *block.Block, c *Connection) message.Message {
	return message.NewRoomsInLobbyMessage(
		s.lobbies[c.LobbyId].Rooms,
	)
}

func GamePlayerInfo(s *Server, b *block.Block, c *Connection) message.Message {
	playerId := block.NewUint32(b)
	player, err := s.Storage().GetPlayer(playerId.Value)
	if err != nil {
		s.Log(c, "Unable to get player %d: %s", playerId.Value, err)
		return nil
	}
	return message.NewGamePlayerInfo(
		block.PlayerInfo{player},
	)
}

func Unknown308c(_ *Server, _ *block.Block, _ *Connection) message.Message {
	// Contains a byte with a 1 in my records
	return message.NewUnknown308cMessage()
}

func Chat(s *Server, b *block.Block, c *Connection) message.Message {
	chatMessage := block.NewChatMessage(b, c.Player.Name)
	s.Log(c, "Received chat message: %v", chatMessage)
	// for now just broadcast the message to everyone
	sendToLobby(s.connections, c.LobbyId, message.NewChatMessage(
		chatMessage,
	))
	return nil
}

func StartGame(stor storage.Storage) {
	fmt.Println("Game Server starting")
	s := NewServer(log.New(os.Stdout, "Game: ", log.LstdFlags), NewGameServerHandler(stor))
	s.Serve(10887)
}
