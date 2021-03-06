package server

import (
	"fmt"

	"log"
	"os"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/data/message"
	"github.com/pinfake/pes6go/data/types"
	"github.com/pinfake/pes6go/storage"
)

type GameServerData struct {
	rooms *types.IdMap
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
	0x4320: JoinRoom,
	0x432a: LeaveRoom,
	0x4345: GetRoomPlayerLinks,
	0x434d: ChangeRoom,
	0x4350: RoomSettings,
	0x4363: Participate,
	0x4400: Chat,
	0x4b00: GetPlayerLink,
}

func NewGameServerHandler(stor storage.Storage) GameServer {
	return GameServer{
		storage: stor,
		data:    GameServerData{types.NewIdMap()},
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

func Participate(s *Server, b *block.Block, c *Connection) message.Message {
	participation := block.NewByte(b)
	room := s.lobbies[c.LobbyId].Rooms.Get(c.Player.RoomId).(*block.Room)
	newParticipation, err := room.ToggleParticipation(c.Player.Id)
	if err != nil {
		panic(err)
	}

	sendToRoom(s.connections, room.Id, message.NewRoomParticipation(
		block.NewRoomParticipation(room),
	), nil)

	return message.NewPlayerParticipateResponse(
		&block.PlayerParticipateResponse{
			Code:          0,
			Selection:     participation.Value,
			Participation: newParticipation,
		},
	)
}

func ChangeRoom(s *Server, b *block.Block, c *Connection) message.Message {
	changeRoom := block.NewCreateRoom(b)
	room := s.lobbies[c.LobbyId].Rooms.Get(c.Player.RoomId).(*block.Room)
	room.HasPassword = changeRoom.HasPassword
	room.Name = changeRoom.Name
	room.HasPassword = changeRoom.HasPassword
	sendToLobby(s.connections, c.LobbyId, message.NewRoomUpdate(room))
	return message.NewChangeRoomResponse(0)
}

// TODO: Prevent creation of rooms with an existing name (in the same lobby)
func CreateRoom(s *Server, b *block.Block, c *Connection) message.Message {
	createRoom := block.NewCreateRoom(b)
	// Should I do this here?, it should be done at joining lobby most probably
	c.Player.ResetRoomData()
	s.Log(c, "Create room: %v", createRoom)
	room := &block.Room{
		Id:          s.Data().(GameServerData).rooms.GetNewId(),
		Type:        1,
		Name:        createRoom.Name,
		HasPassword: createRoom.HasPassword,
		Password:    createRoom.Password,
		Players: []*block.Player{
			c.Player,
		},
	}
	s.lobbies[c.LobbyId].Rooms.Add(room.Id, room)
	c.Player.RoomId = room.Id
	sendToLobby(s.connections, c.LobbyId, message.NewRoomUpdate(room))
	// Maybe just to me?, pes6j says to send this info for every player in the room to me when "entering"
	sendToLobby(s.connections, c.LobbyId, message.NewPlayerUpdate(c.Player))
	//c.writeMessage(message.NewRoomUpdate(room))
	//c.writeMessage(message.NewPlayerUpdate(*c.Player))
	return message.NewCreateRoomResponse()
}

func RoomSettings(s *Server, b *block.Block, c *Connection) message.Message {
	sendToRoom(s.connections, c.Player.RoomId, message.NewReplayBlock(b), c)
	return nil
}

func GetRoomPlayerLinks(s *Server, b *block.Block, c *Connection) message.Message {
	roomId := block.NewUint32(b)
	room := s.lobbies[c.LobbyId].Rooms.Get(roomId.Value).(*block.Room)
	// TODO: Yeah, you're so bright tonight... you forgot about all this
	// TODO: Pes6j says send room update to lobby,
	// TODO: sixserver says 0X4330 which never happended on pes6j, but maybe it helps at some bug i can't tell right now
	//	sendToLobby(s.connections, c.LobbyId, message.NewRoomUpdate(room))

	return message.NewRoomPlayerLinks(
		block.NewRoomPlayerLinks(room),
	)
}

func getRoom(s *Server, c *Connection, roomId uint32) *block.Room {
	return s.lobbies[c.LobbyId].Rooms.Get(roomId).(*block.Room)
}

func JoinRoom(s *Server, b *block.Block, c *Connection) message.Message {
	joinData := block.NewJoinRoomRequest(b)
	room := getRoom(s, c, joinData.Id)
	if room == nil {
		s.Log(c, "Non existing %d room, cannot join", joinData.Id)
		return message.NewJoinRoomResponse(block.GeneralError, 0)
	}

	if room.GetNumPlayers() == 4 {
		return message.NewJoinRoomResponse(block.RoomFull, 0)
	}

	if room.HasPassword == 1 && room.Password != joinData.Password {
		return message.NewJoinRoomResponse(block.BadPassword, 0)
	}
	c.Player.RoomId = joinData.Id
	position := room.AddPlayer(c.Player)

	sendToLobby(s.connections, c.LobbyId, message.NewPlayerUpdate(c.Player))
	sendToLobby(s.connections, c.LobbyId, message.NewRoomUpdate(room))
	// This is what pes6j do, not the other way around, must try this at home.

	c.writeMessage(message.NewJoinRoomResponse(block.Ok, position))

	sendToRoom(s.connections, room.Id, message.NewRoomPlayerLinks(
		block.NewRoomPlayerLinks(room),
	), nil)
	return message.NewJoinRoomResponse(block.Ok, position)

}

func LeaveRoom(s *Server, _ *block.Block, c *Connection) message.Message {
	if c.Player.RoomId != 0 {
		roomId := c.Player.RoomId
		lobby := s.lobbies[c.LobbyId]
		room := lobby.Rooms.Get(c.Player.RoomId).(*block.Room)
		room.RemovePlayer(c.Player)
		sendToLobby(s.connections, c.LobbyId, message.NewPlayerUpdate(c.Player))
		sendToLobby(s.connections, c.LobbyId, message.NewRoomUpdate(room))

		// Not always, just on deco or crashing, if the user is just leaving the
		// room, the other player will generate queries get room link info again,
		// but if crashing or maybe closing the game, that won't happen so i have
		// to make sure the room gets updated.
		sendToRoom(s.connections, room.Id, message.NewRoomPlayerLinks(
			block.NewRoomPlayerLinks(room),
		), nil)
		if !room.HasPlayers() {
			lobby.RemoveRoom(roomId)
			sendToLobby(s.connections, c.LobbyId, message.NewRoomDeleted(room.Id))
		}
		return message.NewLeaveRoomResponse()
	}
	return nil
}

func PlayersInLobby(s *Server, _ *block.Block, c *Connection) message.Message {
	return message.NewPlayersInLobby(
		playersInLobby(s.connections, c.LobbyId),
	)
}

func RoomsInLobby(s *Server, _ *block.Block, c *Connection) message.Message {
	return message.NewRoomsInLobby(
		block.GetRoomsSlice(s.lobbies[c.LobbyId].Rooms),
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
		&block.PlayerExtended{player},
	)
}

func Unknown308c(_ *Server, _ *block.Block, _ *Connection) message.Message {
	// Contains a byte with a 1 in my records
	return message.NewUnknown308c()
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

func GetPlayerLink(s *Server, b *block.Block, _ *Connection) message.Message {
	playerId := block.NewUint32(b)
	targetConn := findByPlayerId(s.connections, playerId.Value)
	if targetConn == nil {
		// Sixservers sends the expected 0x4b01 with 0xff, 0xff, 0xff, 0xff
		return nil
	} else {
		return message.NewPlayerLinkResponse(
			&block.PlayerLink{
				Player: targetConn.Player,
			},
		)
	}
}

func StartGame(stor storage.Storage) {
	fmt.Println("Game Server starting")
	s := NewServer(log.New(os.Stdout, "Game: ", log.LstdFlags), NewGameServerHandler(stor))
	s.Serve(10887)
}
