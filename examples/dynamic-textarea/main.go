package main

import (
	"log"

	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	textarea textarea.Model
}

func initialModel() model { _ = "STUB: not implemented"; return *new(model) }

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) statusView() string { _ = "STUB: not implemented"; return "" }

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }
