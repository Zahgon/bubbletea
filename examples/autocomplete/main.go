package main

import (
	"log"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	gotReposSuccessMsg []repo
	gotReposErrMsg     error
)

type repo struct {
	Name string `json:"name"`
}

const reposURL = "https://api.github.com/orgs/charmbracelet/repos"

func getRepos() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

type model struct {
	textInput textinput.Model
	help      help.Model
	keymap    keymap
}

type keymap struct {
	complete, next, prev, quit key.Binding
}

func (k keymap) ShortHelp() []key.Binding { _ = "STUB: not implemented"; return nil }

func (k keymap) FullHelp() [][]key.Binding { _ = "STUB: not implemented"; return nil }

func initialModel() model { _ = "STUB: not implemented"; return *new(model) }

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Cursor() *tea.Cursor { _ = "STUB: not implemented"; return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func (m model) headerView() string { _ = "STUB: not implemented"; return "" }
func (m model) footerView() string { _ = "STUB: not implemented"; return "" }
