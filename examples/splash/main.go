package main

import (
	"fmt"
	"image/color"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var colors = []color.Color{
	lipgloss.Color("#881177"),
	lipgloss.Color("#aa3355"),
	lipgloss.Color("#cc6666"),
	lipgloss.Color("#ee9944"),
	lipgloss.Color("#eedd00"),
	lipgloss.Color("#99dd55"),
	lipgloss.Color("#44dd88"),
	lipgloss.Color("#22ccbb"),
	lipgloss.Color("#00bbcc"),
	lipgloss.Color("#0099cc"),
	lipgloss.Color("#3366bb"),
	lipgloss.Color("#663399"),
}

type model struct {
	width  int
	height int
	rate   int64
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func (m model) gradient() string { _ = "STUB: not implemented"; return "" }

func getGradientColor(position float64) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func interpolateColors(color1, color2 color.Color, t float64) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

type tickMsg time.Time

func tick() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func main() {
	p := tea.NewProgram(
		model{rate: 90},
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}
}
