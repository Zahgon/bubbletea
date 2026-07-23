package tea

import (
	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/x/ansi"
)

type nilRenderer struct{}

var _ renderer = nilRenderer{}

func (n nilRenderer) start() { _ = "STUB: not implemented"; return }

func (n nilRenderer) clearScreen() { _ = "STUB: not implemented"; return }

func (n nilRenderer) insertAbove(string) error { _ = "STUB: not implemented"; return nil }

func (n nilRenderer) resize(int, int) { _ = "STUB: not implemented"; return }

func (n nilRenderer) setColorProfile(colorprofile.Profile) { _ = "STUB: not implemented"; return }

func (nilRenderer) flush(bool) error { _ = "STUB: not implemented"; return nil }

func (nilRenderer) close() error { _ = "STUB: not implemented"; return nil }

func (nilRenderer) render(View) { _ = "STUB: not implemented"; return }

func (nilRenderer) reset() { _ = "STUB: not implemented"; return }

func (nilRenderer) writeString(string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (n nilRenderer) setSyncdUpdates(bool) { _ = "STUB: not implemented"; return }

func (n nilRenderer) setWidthMethod(ansi.Method) { _ = "STUB: not implemented"; return }

func (n nilRenderer) onMouse(MouseMsg) Cmd { _ = "STUB: not implemented"; return *new(Cmd) }
