package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type model bool

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	p := tea.NewProgram(model(false))
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Oh no:", err)
	}
}
