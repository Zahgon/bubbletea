package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/harmonica"
)

const (
	fps       = 60
	frequency = 7.5
	damping   = 0.15
	asterisk  = "*"
)

func drawEllipse(cb *cellbuffer, xc, yc, rx, ry float64) { _ = "STUB: not implemented"; return }

type cellbuffer struct {
	cells  []string
	stride int
}

func (c *cellbuffer) init(w, h int) {
	if w == 0 {
		return
	}
	c.stride = w
	c.cells = make([]string, w*h)
	c.wipe()
}

func (c cellbuffer) set(x, y int) { _ = "STUB: not implemented"; return }

func (c *cellbuffer) wipe() { _ = "STUB: not implemented"; return }

func (c cellbuffer) width() int { _ = "STUB: not implemented"; return 0 }

func (c cellbuffer) height() int { _ = "STUB: not implemented"; return 0 }

func (c cellbuffer) ready() bool { _ = "STUB: not implemented"; return false }

func (c cellbuffer) String() string { _ = "STUB: not implemented"; return "" }

type frameMsg struct{}

func animate() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

type model struct {
	cells                cellbuffer
	spring               harmonica.Spring
	targetX, targetY     float64
	x, y                 float64
	xVelocity, yVelocity float64
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	m := model{
		spring: harmonica.NewSpring(harmonica.FPS(fps), frequency, damping),
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Uh oh:", err)
		os.Exit(1)
	}
}
