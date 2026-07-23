package tea

import (
	"os"
)

func (p *Program) suspend() { _ = "STUB: not implemented"; return }

func (p *Program) initTerminal() error { _ = "STUB: not implemented"; return nil }

func (p *Program) restoreTerminalState() error { _ = "STUB: not implemented"; return nil }

func (p *Program) restoreInput() error { _ = "STUB: not implemented"; return nil }

func (p *Program) initInputReader(cancel bool) error { _ = "STUB: not implemented"; return nil }

func (p *Program) readLoop() { _ = "STUB: not implemented"; return }

func (p *Program) waitForReadLoop() { _ = "STUB: not implemented"; return }

//nolint:mnd

func (p *Program) checkResize() { _ = "STUB: not implemented"; return }

func OpenTTY() (*os.File, *os.File, error) { _ = "STUB: not implemented"; return nil, nil, nil }
