package main

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type model struct {
	input textinput.Model
	width int
}

var _ tea.Model = model{}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	m := model{}
	m.input = textinput.New()
	m.input.Placeholder = "Enter capability name to request"
	m.input.Focus()

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Uh oh:", err)
		os.Exit(1)
	}
}

func min(a, b int) int { _ = "STUB: not implemented"; return 0 }
