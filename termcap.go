package tea

type requestCapabilityMsg string

func RequestCapability(s string) Cmd { _ = "STUB: not implemented"; return *new(Cmd) }

type CapabilityMsg struct {
	Content string
}

func (c CapabilityMsg) String() string { _ = "STUB: not implemented"; return "" }
