package client

import "github.com/pinfake/pes6go/client/command"

type Config struct {
	ip       string
	port     int
	commands []command.Command
}

type AutomatedClient struct {
	config Config
}

func (au AutomatedClient) run() {
	c := NewClient()
	c.Connect(au.config.ip, au.config.port)
	for _, cmd := range au.config.commands {
		cmd.Execute(c)
	}
}

func NewAutomatedClient(ip string, port int, commands []command.Command) AutomatedClient {
	return AutomatedClient{
		Config{
			ip:       ip,
			port:     port,
			commands: commands,
		},
	}
}
