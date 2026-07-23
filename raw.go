package tea

type RawMsg struct {
	Msg any
}

func Raw(r any) Cmd { _ = "STUB: not implemented"; return *new(Cmd) }
