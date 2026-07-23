package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type editorFinishedMsg struct{ err error }

func openEditor() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

//nolint:gosec

type model struct {
	altscreenActive bool
	err             error
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	m := model{}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
