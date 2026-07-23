package main

import (
	"image/color"
	"log"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type colorType int

const (
	foreground colorType = iota + 1
	background
	cursor
)

func (c colorType) String() string { _ = "STUB: not implemented"; return "" }

type state int

const (
	chooseState state = iota
	inputState
)

type model struct {
	ti          textinput.Model
	choice      colorType
	state       state
	choiceIndex int
	err         error
	fg, bg, cc  color.Color
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	ti := textinput.New()
	ti.Placeholder = "#ff00ff"
	ti.CharLimit = 156
	ti.SetWidth(20)
	ti.SetVirtualCursor(false)
	p := tea.NewProgram(model{
		ti: ti,
	})

	_, err := p.Run()
	if err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
