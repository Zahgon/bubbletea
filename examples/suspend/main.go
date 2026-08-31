package main

import (
	"errors"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	quitting   bool
	suspending bool
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Println("Error running program:", err)
		if errors.Is(err, tea.ErrInterrupted) {
			os.Exit(130)
		}
		os.Exit(1)
	}
}
