package main

import (
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
)

const debounceDuration = time.Second

type exitMsg int

type model struct {
	tag int
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Println("uh oh:", err)
		os.Exit(1)
	}
}
