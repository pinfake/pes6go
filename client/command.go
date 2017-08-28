package client

type Command struct {
	name string
	Data map[string]interface{}
	CommandHandler
}

type CommandHandler interface {
	Execute(*Client)
}
