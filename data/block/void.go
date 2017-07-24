package block

type Void struct {
}

type VoidInternal struct {
}

func (info Void) buildInternal() PieceInternal {
	return VoidInternal{}
}
