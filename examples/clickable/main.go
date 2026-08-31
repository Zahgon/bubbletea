package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/segmentio/ksuid"
)

type LayerHitMsg struct {
	ID    string
	Mouse tea.MouseMsg
}

const maxDialogs = 999

var (
	bgTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("239")).
			Padding(1, 2)

	bgWhitespace = []lipgloss.WhitespaceOption{
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("238"))),
	}

	dialogWordStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E7E1CC"))

	dialogStyle = dialogWordStyle.
			Width(36).
			Height(8).
			Padding(1, 3).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD"))

	hoveredDialogStyle = dialogStyle.
				BorderForeground(lipgloss.Color("#F25D94"))

	specialWordLightColor = lipgloss.Color("#43BF6D")
	specialWordDarkColor  = lipgloss.Color("#73F59F")

	buttonStyle = lipgloss.NewStyle().
			Padding(0, 3).
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#6124DF"))

	hoveredButtonStyle = buttonStyle.
				Background(lipgloss.Color("#FF5F87"))
)

type model struct {
	specialWordStyle lipgloss.Style
	width, height    int
	dialogs          []dialog
	mouseDown        bool
	pressID          string
	dragID           string
	dragOffsetX      int
	dragOffsetY      int
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

func (m *model) newDialog(x, y int) (d dialog) { _ = "STUB: not implemented"; return *new(dialog) }

func (m model) removeDialog(index int) []dialog { _ = "STUB: not implemented"; return nil }

type dialog struct {
	specialWordStyle *lipgloss.Style
	id               string
	buttonID         string
	x, y             int
	text             string
	hovering         bool
	hoveringButton   bool
}

func (d dialog) buttonView() string { _ = "STUB: not implemented"; return "" }

func (d dialog) windowView() string { _ = "STUB: not implemented"; return "" }

func (d dialog) view() *lipgloss.Layer { _ = "STUB: not implemented"; return nil }

func main() {
	ksuid.SetRand(ksuid.FastRander)

	path := os.Getenv("TEA_LOGFILE")
	if path != "" {
		f, err := tea.LogToFile(path, "layers")
		if err != nil {
			fmt.Println("could not open logfile:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Println("Error while running program:", err)
		os.Exit(1)
	}
}

func clamp(n, min, max int) int { _ = "STUB: not implemented"; return 0 }
