package main

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var whiteFg = lipgloss.NewStyle().Foreground(lipgloss.White)

type model struct {
	screenBuf   []int
	width       int
	height      int
	firePalette []int
	startTime   time.Time
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func (m *model) spreadFire() { _ = "STUB: not implemented"; return }

func (m *model) spreadPixel(idx int) { _ = "STUB: not implemented"; return }

type tickMsg time.Time

func tick() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func initialModel() model { _ = "STUB: not implemented"; return *new(model) }

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}
}
