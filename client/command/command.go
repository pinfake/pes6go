package command

import (
	"github.com/pinfake/pes6go/client"
)

type Command struct {
	name string
	data map[string]interface{}
}

type CommandHandler interface {
	execute(*client.Client)
}
