package block

type MenuServers struct {
	MenuServers []MenuServer
}

type MenuServer struct {
	Type       byte
	Name       string
	NumClients uint16
}

type MenuServerInternal struct {
	Type       byte
	Name       [32]byte
	NumClients uint16
}

type MenuServersInternal struct {
	NumServers          uint16
	MenuServersInternal []MenuServerInternal
}

func (info MenuServers) buildInternal() PieceInternal {
	internals := MenuServersInternal{
		NumServers: uint16(len(info.MenuServers)),
	}

	for _, server := range info.MenuServers {
		var internal MenuServerInternal
		internal.Type = server.Type
		copy(internal.Name[:], server.Name)
		internal.NumClients = server.NumClients
		internals.MenuServersInternal = append(internals.MenuServersInternal, internal)
	}

	return internals
}
