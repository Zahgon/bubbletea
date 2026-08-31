package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
)

func main() {
	p := tea.NewProgram(model{
		focused:   true,
		reporting: true,
	})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	focused   bool
	reporting bool
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }
