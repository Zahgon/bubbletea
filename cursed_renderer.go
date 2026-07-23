package tea

import (
	"bytes"
	"io"
	"sync"

	"github.com/charmbracelet/colorprofile"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
)

type cursedRenderer struct {
	w             io.Writer
	buf           bytes.Buffer
	scr           *uv.TerminalRenderer
	cellbuf       uv.ScreenBuffer
	lastView      *View
	env           []string
	term          string
	width, height int
	mu            sync.Mutex
	profile       colorprofile.Profile
	logger        uv.Logger
	view          View
	hardTabs      bool
	backspace     bool
	mapnl         bool
	syncdUpdates  bool
	starting      bool
}

var _ renderer = &cursedRenderer{}

func newCursedRenderer(w io.Writer, env []string, width, height int) (s *cursedRenderer) {
	_ = "STUB: not implemented"
	return nil
}

func (s *cursedRenderer) setLogger(logger uv.Logger) { _ = "STUB: not implemented"; return }

func (s *cursedRenderer) setOptimizations(hardTabs, backspace, mapnl bool) {
	_ = "STUB: not implemented"
	return
}

func (s *cursedRenderer) start() { _ = "STUB: not implemented"; return }

func (s *cursedRenderer) close() (err error) { _ = "STUB: not implemented"; return nil }

//nolint:nestif

func (s *cursedRenderer) writeString(str string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:wrapcheck

func (s *cursedRenderer) flush(closing bool) error { _ = "STUB: not implemented"; return nil }

func (s *cursedRenderer) render(v View) { _ = "STUB: not implemented"; return }

func (s *cursedRenderer) reset() { _ = "STUB: not implemented"; return }

func reset(s *cursedRenderer) { _ = "STUB: not implemented"; return }

func (s *cursedRenderer) setColorProfile(p colorprofile.Profile) { _ = "STUB: not implemented"; return }

func (s *cursedRenderer) resize(w, h int) { _ = "STUB: not implemented"; return }

func (s *cursedRenderer) clearScreen() { _ = "STUB: not implemented"; return }

func enableAltScreen(s *cursedRenderer, enable bool, write bool) { _ = "STUB: not implemented"; return }

func enterAltScreen(s *cursedRenderer, write bool) { _ = "STUB: not implemented"; return }

func exitAltScreen(s *cursedRenderer, write bool) { _ = "STUB: not implemented"; return }

func enableTextCursor(s *cursedRenderer, enable bool) { _ = "STUB: not implemented"; return }

func (s *cursedRenderer) setSyncdUpdates(syncd bool) { _ = "STUB: not implemented"; return }

func (s *cursedRenderer) setWidthMethod(method ansi.Method) { _ = "STUB: not implemented"; return }

func (s *cursedRenderer) insertAbove(str string) error { _ = "STUB: not implemented"; return nil }

func (s *cursedRenderer) onMouse(m MouseMsg) Cmd { _ = "STUB: not implemented"; return *new(Cmd) }

func setProgressBar(s *cursedRenderer, pb *ProgressBar) { _ = "STUB: not implemented"; return }

func viewEquals(a, b *View) bool { _ = "STUB: not implemented"; return false }

func keyboardEnhancementsFlags(ke KeyboardEnhancements) int { _ = "STUB: not implemented"; return 0 }
