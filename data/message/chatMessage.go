package message

import (
	"github.com/pinfake/pes6go/data/block"
)

type ChatMessage struct {
	*block.ChatMessage
}

func (data ChatMessage) GetBlocks() []*block.Block {
	return block.GetBlocks(0x4402, data.ChatMessage)
}

func NewChatMessage(info *block.ChatMessage) ChatMessage {
	return ChatMessage{info}
}
