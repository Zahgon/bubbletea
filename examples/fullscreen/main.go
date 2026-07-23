package main

import (
	"log"
	"time"

	tea "charm.land/bubbletea/v2"
)

type model int

type tickMsg time.Time

func main() {
	p := tea.NewProgram(model(5))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func tick() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }
