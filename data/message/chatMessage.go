package message

import (
	"reflect"

	"github.com/pinfake/pes6go/data/block"
)

type ChatMessage struct {
	ChatMessage []block.Piece
}

func (r ChatMessage) GetBlocks() []*block.Block {
	var blocks []*block.Block

	blocks = append(blocks, block.GetBlocks(0x4402, r.ChatMessage)...)

	return blocks
}

func NewChatMessage(info block.ChatMessage) ChatMessage {
	return ChatMessage{
		ChatMessage: block.GetPieces(reflect.ValueOf(info)),
	}
}
