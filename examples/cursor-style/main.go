package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	cursor tea.Cursor
	blink  bool
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func (m model) describeCursor() string { _ = "STUB: not implemented"; return "" }

func main() {
	p := tea.NewProgram(model{blink: true})
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
