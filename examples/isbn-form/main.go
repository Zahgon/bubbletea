package main

import (
	"log"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/exp/charmtone"
)

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color(charmtone.Tang.Hex()))
	continueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(charmtone.Anchovy.Hex()))
	validStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color(charmtone.Guac.Hex()))
	errStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color(charmtone.Cherry.Hex()))
)

type model struct {
	isbnInput    textinput.Model
	titleInput   textinput.Model
	focusedInput int
	err          error
}

func (m model) canFindBook() bool { _ = "STUB: not implemented"; return false }

func isbn13Validator(s string) error { _ = "STUB: not implemented"; return nil }

var bannedTitleWords = []string{
	"very",
	"bad",
	"words",
	"that",
	"should",
	"not",
	"appear",
	"in",
	"book",
	"titles",
}

func bookTitleValidator(s string) error { _ = "STUB: not implemented"; return nil }

func initialModel() model { _ = "STUB: not implemented"; return *new(model) }

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }
