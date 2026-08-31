package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

var choices = []string{"Taro", "Coffee", "Lychee"}

type model struct {
	cursor int
	choice string
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func main() {
	p := tea.NewProgram(model{})

	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	if m, ok := m.(model); ok && m.choice != "" {
		fmt.Printf("\n---\nYou chose %s!\n", m.choice)
	}
}
