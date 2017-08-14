package block

type Byte struct {
	Value byte
}

type ByteInternal struct {
	Value byte
}

func (info Byte) buildInternal() PieceInternal {
	var internal ByteInternal
	internal.Value = info.Value
	return internal
}

func NewByte(b *Block) Byte {
	return Byte{
		Value: b.Body.GetBytes()[0],
	}
}
