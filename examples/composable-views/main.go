package main

import (
	"log"
	"time"

	"charm.land/bubbles/v2/spinner"
	"charm.land/bubbles/v2/timer"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type sessionState uint

const (
	defaultTime              = time.Minute
	timerView   sessionState = iota
	spinnerView
)

var (
	spinners = []spinner.Spinner{
		spinner.Line,
		spinner.Dot,
		spinner.MiniDot,
		spinner.Jump,
		spinner.Pulse,
		spinner.Points,
		spinner.Globe,
		spinner.Moon,
		spinner.Monkey,
	}
	modelStyle = lipgloss.NewStyle().
			Width(15).
			Height(5).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.HiddenBorder())
	focusedModelStyle = lipgloss.NewStyle().
				Width(15).
				Height(5).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("69"))
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type mainModel struct {
	state   sessionState
	timer   timer.Model
	spinner spinner.Model
	index   int
}

func newModel(timeout time.Duration) mainModel { _ = "STUB: not implemented"; return *new(mainModel) }

func (m mainModel) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m mainModel) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func (m mainModel) currentFocusedModel() string { _ = "STUB: not implemented"; return "" }

func (m *mainModel) Next() { _ = "STUB: not implemented"; return }

func (m *mainModel) resetSpinner() { _ = "STUB: not implemented"; return }

func main() {
	p := tea.NewProgram(newModel(defaultTime))

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
