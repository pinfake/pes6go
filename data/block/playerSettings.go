package block

type PlayerSettingsHeader struct {
	PlayerId uint32
}

type PlayerSettings struct {
	Settings []byte
}

type PlayerSettingsHeaderInternal struct {
	Zero     uint32
	PlayerId uint32
}

type PlayerSettingsInternal struct {
	Settings [1300]byte
}

func (info PlayerSettings) buildInternal() PieceInternal {
	var internal PlayerSettingsInternal
	copy(internal.Settings[:], info.Settings)

	return internal
}

func (info PlayerSettingsHeader) buildInternal() PieceInternal {
	var internal PlayerSettingsHeaderInternal
	internal.Zero = 0
	internal.PlayerId = info.PlayerId

	return internal
}
