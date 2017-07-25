package block

import "encoding/binary"

type ChatMessage struct {
	Place      byte // lobby = 0, or room = 1
	Channel    byte // 1 = to lobby, 2, 3 = private, 4 = group, 5 = playing, 7 = in-game, 8 = room, 9 = team
	PlayerId   uint32
	PlayerName string
	Message    string
}

type ChatMessageInternal struct {
	Place      byte
	Channel    byte
	Unknown1   [4]byte
	PlayerId   uint32
	PlayerName [48]byte
	Message    [48]byte
}

func (info ChatMessage) buildInternal() PieceInternal {
	var internal ChatMessageInternal
	internal.Place = info.Place
	internal.Channel = info.Channel
	internal.PlayerId = info.PlayerId
	copy(internal.PlayerName[:], info.PlayerName)
	internal.Unknown1 = [4]byte{0xff, 0xff, 0xff, 0xff}
	copy(internal.Message[:], info.Message)
	return internal
}

func NewChatMessage(b *Block, playerName string) ChatMessage {
	body := b.Body.GetBytes()
	return ChatMessage{
		Place:      body[0],
		Channel:    body[1],
		PlayerName: playerName,
		PlayerId:   binary.BigEndian.Uint32(body[6:10]),
		Message:    string(body[10:]),
	}
}
