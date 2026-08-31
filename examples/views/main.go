package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"
	dotChar           = " • "
)

var (
	keywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	ticksStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("79"))
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	progressEmpty = subtleStyle.Render(progressEmptyChar)
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)

	ramp = makeRampStyles("#B14FFF", "#00FFA3", progressBarWidth)
)

func main() {
	initialModel := model{0, false, 10, 0, 0, false, false}
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}

type (
	tickMsg  struct{}
	frameMsg struct{}
)

func tick() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func frame() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

type model struct {
	Choice   int
	Chosen   bool
	Ticks    int
	Frames   int
	Progress float64
	Loaded   bool
	Quitting bool
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func updateChoices(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func updateChosen(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func choicesView(m model) string { _ = "STUB: not implemented"; return "" }

func chosenView(m model) string { _ = "STUB: not implemented"; return "" }

func checkbox(label string, checked bool) string { _ = "STUB: not implemented"; return "" }

func progressbar(percent float64) string { _ = "STUB: not implemented"; return "" }

func makeRampStyles(colorA, colorB string, steps float64) (s []lipgloss.Style) {
	_ = "STUB: not implemented"
	return nil
}

func colorToHex(c colorful.Color) string { _ = "STUB: not implemented"; return "" }

func colorFloatToHex(f float64) (s string) { _ = "STUB: not implemented"; return "" }
