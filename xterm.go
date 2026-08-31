package tea

type TerminalVersionMsg struct {
	Name string
}

func (t TerminalVersionMsg) String() string { _ = "STUB: not implemented"; return "" }

type terminalVersion struct{}

func RequestTerminalVersion() Msg { _ = "STUB: not implemented"; return *new(Msg) }
