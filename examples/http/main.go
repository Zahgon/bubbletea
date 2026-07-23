package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
)

const url = "https://charm.sh/"

type model struct {
	status int
	err    error
}

type statusMsg int

type errMsg struct{ error }

func (e errMsg) Error() string { _ = "STUB: not implemented"; return "" }

func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func checkServer() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }
