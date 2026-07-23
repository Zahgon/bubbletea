package tea

import (
	"bytes"
	"context"
	"errors"
	"image/color"
	"io"
	"sync"
	"time"

	"github.com/charmbracelet/colorprofile"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/term"
	"github.com/muesli/cancelreader"
)

var ErrProgramPanic = errors.New("program experienced a panic")

var ErrProgramKilled = errors.New("program was killed")

var ErrInterrupted = errors.New("program was interrupted")

type Msg = uv.Event

type Model interface {
	Init() Cmd

	Update(Msg) (Model, Cmd)

	View() View
}

func NewView(s string) View { _ = "STUB: not implemented"; return *new(View) }

type View struct {
	Content string

	OnMouse func(msg MouseMsg) Cmd

	Cursor *Cursor

	BackgroundColor color.Color

	ForegroundColor color.Color

	WindowTitle string

	ProgressBar *ProgressBar

	AltScreen bool

	ReportFocus bool

	DisableBracketedPasteMode bool

	MouseMode MouseMode

	KeyboardEnhancements KeyboardEnhancements
}

type KeyboardEnhancements struct {
	ReportEventTypes bool

	ReportAlternateKeys bool

	ReportAllKeysAsEscapeCodes bool

	ReportAssociatedText bool
}

func (v *View) SetContent(s string) { _ = "STUB: not implemented"; return }

type MouseMode int

const (
	MouseModeNone MouseMode = iota

	MouseModeCellMotion

	MouseModeAllMotion
)

type ProgressBarState int

const (
	ProgressBarNone ProgressBarState = iota
	ProgressBarDefault
	ProgressBarError
	ProgressBarIndeterminate
	ProgressBarWarning
)

func (s ProgressBarState) String() string { _ = "STUB: not implemented"; return "" }

type ProgressBar struct {
	State ProgressBarState

	Value int
}

func NewProgressBar(state ProgressBarState, value int) *ProgressBar {
	_ = "STUB: not implemented"
	return nil
}

type Cursor struct {
	Position

	Color color.Color

	Shape CursorShape

	Blink bool
}

func NewCursor(x, y int) *Cursor { _ = "STUB: not implemented"; return nil }

type Cmd func() Msg

type channelHandlers struct {
	handlers []chan struct{}
	mu       sync.RWMutex
}

func (h *channelHandlers) add(ch chan struct{}) { _ = "STUB: not implemented"; return }

func (h *channelHandlers) shutdown() { _ = "STUB: not implemented"; return }

type Program struct {
	disableInput bool

	disableSignalHandler bool

	disableCatchPanics bool

	filter func(Model, Msg) Msg

	fps int

	initialModel Model

	disableRenderer bool

	handlers channelHandlers

	ctx    context.Context
	cancel context.CancelFunc

	externalCtx context.Context

	msgs         chan Msg
	errs         chan error
	finished     chan struct{}
	shutdownOnce sync.Once

	profile *colorprofile.Profile

	output    io.Writer
	outputBuf bytes.Buffer

	ttyOutput           term.File
	previousOutputState *term.State
	renderer            renderer

	environ uv.Environ

	logger uv.Logger

	input io.Reader

	ttyInput              term.File
	previousTtyInputState *term.State
	cancelReader          cancelreader.CancelReader
	inputScanner          *uv.TerminalReader
	readLoopDone          chan struct{}

	ignoreSignals uint32

	ticker *time.Ticker

	once sync.Once

	rendererDone chan struct{}

	width, height int

	useHardTabs bool

	useBackspace bool

	mu sync.Mutex
}

func Quit() Msg { _ = "STUB: not implemented"; return *new(Msg) }

type QuitMsg struct{}

func Suspend() Msg { _ = "STUB: not implemented"; return *new(Msg) }

type SuspendMsg struct{}

type ResumeMsg struct{}

type InterruptMsg struct{}

func Interrupt() Msg { _ = "STUB: not implemented"; return *new(Msg) }

func NewProgram(model Model, opts ...ProgramOption) *Program { _ = "STUB: not implemented"; return nil }

func (p *Program) handleSignals() chan struct{} { _ = "STUB: not implemented"; return nil }

func (p *Program) handleResize() chan struct{} { _ = "STUB: not implemented"; return nil }

func (p *Program) handleCommands(cmds chan Cmd) chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (p *Program) eventLoop(model Model, cmds chan Cmd) (Model, error) {
	_ = "STUB: not implemented"
	return *new(Model), nil
}

//nolint:errcheck,gosec

func (p *Program) render(model Model) { _ = "STUB: not implemented"; return }

func (p *Program) execSequenceMsg(msg sequenceMsg) { _ = "STUB: not implemented"; return }

func (p *Program) execBatchMsg(msg BatchMsg) { _ = "STUB: not implemented"; return }

func shouldQuerySynchronizedOutput(environ uv.Environ) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Program) Run() (returnModel Model, returnErr error) {
	_ = "STUB: not implemented"
	return *new(Model), nil
}

func (p *Program) Send(msg Msg) { _ = "STUB: not implemented"; return }

func (p *Program) Quit() { _ = "STUB: not implemented"; return }

func (p *Program) Kill() { _ = "STUB: not implemented"; return }

func (p *Program) Wait() { _ = "STUB: not implemented"; return }

func (p *Program) execute(seq string) { _ = "STUB: not implemented"; return }

func (p *Program) flush() error { _ = "STUB: not implemented"; return nil }

func (p *Program) shutdown(kill bool) { _ = "STUB: not implemented"; return }

func (p *Program) recoverFromPanic(r interface{}) { _ = "STUB: not implemented"; return }

//nolint:errcheck
//nolint:errcheck
//nolint:errcheck
//nolint:errcheck

func (p *Program) recoverFromGoPanic(r interface{}) { _ = "STUB: not implemented"; return }

//nolint:errcheck
//nolint:errcheck
//nolint:errcheck
//nolint:errcheck

func (p *Program) ReleaseTerminal() error { _ = "STUB: not implemented"; return nil }

func (p *Program) releaseTerminal(reset bool) error { _ = "STUB: not implemented"; return nil }

func (p *Program) RestoreTerminal() error { _ = "STUB: not implemented"; return nil }

func (p *Program) Println(args ...any) { _ = "STUB: not implemented"; return }

func (p *Program) Printf(template string, args ...any) { _ = "STUB: not implemented"; return }

func (p *Program) startRenderer() { _ = "STUB: not implemented"; return }

func (p *Program) stopRenderer(kill bool) { _ = "STUB: not implemented"; return }
