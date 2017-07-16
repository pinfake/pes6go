package block

type PlayerSettingsHeader struct {
	PlayerId uint32
}

type PlayerSettings struct {
	Settings []byte
}

type PlayerSettingsHeaderInternal struct {
	zero     uint32
	playerId uint32
}

type PlayerSettingsInternal struct {
	settings [1300]byte
}

func (info PlayerSettings) buildInternal() PieceInternal {
	var internal PlayerSettingsInternal
	copy(internal.settings[:], info.Settings)

	return internal
}

func (info PlayerSettingsHeader) buildInternal() PieceInternal {
	var internal PlayerSettingsHeaderInternal
	internal.zero = 0
	internal.playerId = info.PlayerId

	return internal
}
