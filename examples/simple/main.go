package main

import (
	"log"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
)

func main() {

	logfilePath := os.Getenv("BUBBLETEA_LOG")
	if logfilePath != "" {
		if _, err := tea.LogToFile(logfilePath, "simple"); err != nil {
			log.Fatal(err)
		}
	}

	p := tea.NewProgram(model(5))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model int

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

type tickMsg time.Time

func tick() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }
