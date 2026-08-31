package main

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
)

const (
	eyeWidth   = 15
	eyeHeight  = 12
	eyeSpacing = 40

	blinkFrames = 20
	openTimeMin = 1000
	openTimeMax = 4000
)

const (
	eyeChar = "●"
	bgChar  = " "
)

type model struct {
	width        int
	height       int
	eyePositions [2]int
	eyeY         int
	isBlinking   bool
	blinkState   int
	lastBlink    time.Time
	openTime     time.Duration
}

type tickMsg time.Time

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
	}
}

func initialModel() model { _ = "STUB: not implemented"; return *new(model) }

func (m *model) updateEyePositions() { _ = "STUB: not implemented"; return }

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func tickCmd() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func drawEllipse(canvas [][]string, x0, y0, rx, ry int) { _ = "STUB: not implemented"; return }
