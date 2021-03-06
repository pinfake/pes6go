package block

type AccountingInit struct {
}

type AccountingInitInternal struct {
	Unknown [16]byte
}

func (info AccountingInit) buildInternal() PieceInternal {
	var internal AccountingInitInternal

	internal.Unknown = [16]byte{
		// THIS AFFECTS HOW A PASSWORD IS BLOWFISHED, IF ITS ZEROS
		// THEN IT BEHAVES LIKE SIXSERVER AND I CAN USE ITS BLOWFISH
		// MAYBE ITS A SALT OR SOMETHING SIMILAR.
		//0x38, 0x2b, 0x46, 0x47, 0x02, 0x4b, 0x2f, 0x68, // mine
		//0x56, 0x28, 0x3f, 0x53, 0x10, 0x87, 0x32, 0xa0,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // sixserver
	}

	return internal
}
