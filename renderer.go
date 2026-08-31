package tea

import (
	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/x/ansi"
)

const (
	defaultFPS = 60
	maxFPS     = 120
)

type renderer interface {
	start()

	close() error

	render(View)

	flush(closing bool) error

	reset()

	insertAbove(string) error

	setSyncdUpdates(bool)

	setWidthMethod(ansi.Method)

	resize(int, int)

	setColorProfile(colorprofile.Profile)

	clearScreen()

	writeString(string) (int, error)

	onMouse(MouseMsg) Cmd
}

type printLineMessage struct {
	messageBody string
}

func Println(args ...any) Cmd { _ = "STUB: not implemented"; return *new(Cmd) }

func Printf(template string, args ...any) Cmd { _ = "STUB: not implemented"; return *new(Cmd) }

func encodeCursorStyle(style CursorShape, blink bool) int { _ = "STUB: not implemented"; return 0 }

//nolint:mnd
