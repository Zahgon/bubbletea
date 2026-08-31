package tea

type ClipboardMsg struct {
	Content   string
	Selection byte
}

func (e ClipboardMsg) Clipboard() byte { _ = "STUB: not implemented"; return 0 }

func (e ClipboardMsg) String() string { _ = "STUB: not implemented"; return "" }

type setClipboardMsg string

func SetClipboard(s string) Cmd { _ = "STUB: not implemented"; return *new(Cmd) }

type readClipboardMsg struct{}

func ReadClipboard() Msg { _ = "STUB: not implemented"; return *new(Msg) }

type setPrimaryClipboardMsg string

func SetPrimaryClipboard(s string) Cmd { _ = "STUB: not implemented"; return *new(Cmd) }

type readPrimaryClipboardMsg struct{}

func ReadPrimaryClipboard() Msg { _ = "STUB: not implemented"; return *new(Msg) }
