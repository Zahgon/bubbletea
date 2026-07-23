package main

import (
	"fmt"
	"os"
	"time"

	"charm.land/bubbles/v2/progress"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const (
	padding  = 2
	maxWidth = 80
)

var (
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
	yellow    = lipgloss.Color("#FDFF8C")
	pink      = lipgloss.Color("#FF7CCB")
)

func main() {
	prog := progress.New(progress.WithScaled(true), progress.WithColors(pink, yellow))

	if _, err := tea.NewProgram(model{progress: prog}).Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}

type tickMsg time.Time

type model struct {
	percent  float64
	progress progress.Model
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func tickCmd() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }
