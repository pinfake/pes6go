package client

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
		cmd.Execute(c)
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
