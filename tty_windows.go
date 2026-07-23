//go:build windows
// +build windows

package tea

func (p *Program) initInput() (err error) { _ = "STUB: not implemented"; return nil }

//nolint:godox

//nolint:nakedret

const suspendSupported = false

func suspendProcess() { _ = "STUB: not implemented"; return }
