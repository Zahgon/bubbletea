package main

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
)

type responseMsg struct{}

func listenForActivity(sub chan struct{}) tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func waitForActivity(sub chan struct{}) tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

type model struct {
	sub       chan struct{}
	responses int
	spinner   spinner.Model
	quitting  bool
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	p := tea.NewProgram(model{
		sub:     make(chan struct{}),
		spinner: spinner.New(),
	})

	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}
}
