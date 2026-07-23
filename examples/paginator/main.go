package main

import (
	"log"

	"charm.land/bubbles/v2/paginator"
	"charm.land/lipgloss/v2"

	tea "charm.land/bubbletea/v2"
)

type styles struct {
	activeDot   lipgloss.Style
	inactiveDot lipgloss.Style
}

func newStyles(bgIsDark bool) (s styles) { _ = "STUB: not implemented"; return *new(styles) }

type model struct {
	items     []string
	paginator paginator.Model
}

func newModel() model { _ = "STUB: not implemented"; return *new(model) }

func (m *model) updateStyles(isDark bool) { _ = "STUB: not implemented"; return }

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	p := tea.NewProgram(newModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
