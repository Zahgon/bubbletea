package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var body = lipgloss.NewStyle().Padding(1, 2)

type model struct {
	value int
	width int
	state tea.ProgressBarState
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	p := tea.NewProgram(model{value: 50, state: tea.ProgressBarIndeterminate})
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
