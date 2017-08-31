package client

import "fmt"

type Config struct {
	ip       string
	port     int
	commands []CommandHandler
}

type AutomatedClient struct {
	config Config
}

func (au AutomatedClient) Run() {
	c := NewClient()
	c.Connect(au.config.ip, au.config.port)
	for _, cmd := range au.config.commands {
		fmt.Printf("Sending command: %v\n", cmd)
		cmd.Execute(c)
		c.Read()
	}
}

func NewAutomatedClient(ip string, port int, commands []CommandHandler) AutomatedClient {
	return AutomatedClient{
		Config{
			ip:       ip,
			port:     port,
			commands: commands,
		},
	}
}
