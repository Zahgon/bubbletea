package tea

import "github.com/charmbracelet/x/ansi"

type WindowSizeMsg struct {
	Width  int
	Height int
}

func ClearScreen() Msg { _ = "STUB: not implemented"; return *new(Msg) }

type clearScreenMsg struct{}

type ModeReportMsg struct {
	Mode ansi.Mode

	Value ansi.ModeSetting
}
