package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type styles struct {
	ui lipgloss.Style
}

type model struct {
	supportsDisambiguation bool
	supportsEventTypes     bool
	styles                 styles
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func (m *model) updateStyles(isDark bool) { _ = "STUB: not implemented"; return }

func initialModel() model { _ = "STUB: not implemented"; return *new(model) }

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Urgh: %v\n", err)
		os.Exit(1)
	}
}
