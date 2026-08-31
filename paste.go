package tea

type PasteMsg struct {
	Content string
}

func (p PasteMsg) String() string { _ = "STUB: not implemented"; return "" }

type PasteStartMsg struct{}

type PasteEndMsg struct{}
