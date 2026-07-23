//go:build darwin || linux || solaris || aix
// +build darwin linux solaris aix

package tea

import (
	"github.com/charmbracelet/x/term"
)

func (p *Program) checkOptimizedMovements(s *term.State) { _ = "STUB: not implemented"; return }
