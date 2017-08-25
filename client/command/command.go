package command

import (
	"github.com/pinfake/pes6go/client"
	"github.com/pinfake/pes6go/data/message"
)

type Command struct {
	name string
	data map[string]interface{}
}

type CommandHandler interface {
	execute(*client.Client) message.Message
}
