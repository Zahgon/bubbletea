package tea

type KeyboardEnhancementsMsg struct {
	Flags int
}

func (k KeyboardEnhancementsMsg) SupportsKeyDisambiguation() bool {
	_ = "STUB: not implemented"
	return false
}

func (k KeyboardEnhancementsMsg) SupportsEventTypes() bool { _ = "STUB: not implemented"; return false }

func (k KeyboardEnhancementsMsg) SupportsAlternateKeys() bool {
	_ = "STUB: not implemented"
	return false
}

func (k KeyboardEnhancementsMsg) SupportsAllKeysAsEscapeCodes() bool {
	_ = "STUB: not implemented"
	return false
}

func (k KeyboardEnhancementsMsg) SupportsAssociatedText() bool {
	_ = "STUB: not implemented"
	return false
}
