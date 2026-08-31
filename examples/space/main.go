package main

import (
	"fmt"
	"image/color"
	"os"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	colors     [][]color.Color
	lastWidth  int
	lastHeight int
	frameCount int
	width      int
	height     int
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func tickCmd() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

type tickMsg struct{}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m *model) setupColors() { _ = "STUB: not implemented"; return }

func clamp(value, min, max float64) float64 { _ = "STUB: not implemented"; return 0 }

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	p := tea.NewProgram(model{}, tea.WithFPS(120))

	_, err := p.Run()
	if err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}
