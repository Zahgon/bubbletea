package main

import (
	"image/color"
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/lucasb-eyer/go-colorful"
)

var myFancyColor color.Color

type model struct{}

var _ tea.Model = model{}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	myFancyColor, _ = colorful.Hex("#6b50ff")

	p := tea.NewProgram(model{}, tea.WithColorProfile(colorprofile.TrueColor))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
