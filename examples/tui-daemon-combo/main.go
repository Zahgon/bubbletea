package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/term"
)

var (
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
	mainStyle = lipgloss.NewStyle().MarginLeft(1)
)

func main() {
	var (
		daemonMode bool
		showHelp   bool
	)

	flag.BoolVar(&daemonMode, "d", false, "run as a daemon")
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.Parse()

	if showHelp {
		flag.Usage()
		os.Exit(0)
	}

	opts := []tea.ProgramOption{}
	if daemonMode || !term.IsTerminal(os.Stdout.Fd()) {

		opts = append(opts, tea.WithoutRenderer())
	} else {

		log.SetOutput(io.Discard)
	}

	p := tea.NewProgram(newModel(), opts...)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting Bubble Tea program:", err)
		os.Exit(1)
	}
}

type result struct {
	duration time.Duration
	emoji    string
}

type model struct {
	spinner  spinner.Model
	results  []result
	quitting bool
}

func newModel() model { _ = "STUB: not implemented"; return *new(model) }

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

type processFinishedMsg time.Duration

func runPretendProcess() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func randomEmoji() string { _ = "STUB: not implemented"; return "" }
