package main

import (
	"fmt"
	"io"
	"os"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const listHeight = 14

type styles struct {
	title        lipgloss.Style
	item         lipgloss.Style
	selectedItem lipgloss.Style
	pagination   lipgloss.Style
	help         lipgloss.Style
	quitText     lipgloss.Style
}

func newStyles(darkBG bool) styles { _ = "STUB: not implemented"; return *new(styles) }

type item string

func (i item) FilterValue() string { _ = "STUB: not implemented"; return "" }

type itemDelegate struct {
	styles *styles
}

func (d itemDelegate) Height() int  { _ = "STUB: not implemented"; return 0 }
func (d itemDelegate) Spacing() int { _ = "STUB: not implemented"; return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	_ = "STUB: not implemented"
	return *new(tea.Cmd)
}
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	_ = "STUB: not implemented"
	return
}

type model struct {
	list     list.Model
	choice   string
	styles   styles
	quitting bool
}

func initialModel() model { _ = "STUB: not implemented"; return *new(model) }

func (m *model) updateStyles(isDark bool) { _ = "STUB: not implemented"; return }

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
