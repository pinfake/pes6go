package block

type MenuServers struct {
	MenuServers [1]MenuServer
}

type MenuServer struct {
	Stype      byte
	Name       string
	NumClients uint16
}

type MenuServerInternal struct {
	stype      byte
	name       [32]byte
	numClients uint16
}

type MenuServersInternal struct {
	zero                uint16
	MenuServersInternal [1]MenuServerInternal
}

func (info MenuServers) buildInternal() PieceInternal {
	internals := MenuServersInternal{
		zero: 1,
	}

	for i, server := range info.MenuServers {
		var internal MenuServerInternal
		internal.stype = server.Stype
		copy(internal.name[:], server.Name)
		internal.numClients = server.NumClients
		internals.MenuServersInternal[i] = internal
	}

	return internals
}
