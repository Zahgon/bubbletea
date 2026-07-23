//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || aix || zos
// +build darwin dragonfly freebsd linux netbsd openbsd solaris aix zos

package tea

func (p *Program) initInput() (err error) { _ = "STUB: not implemented"; return nil }

const suspendSupported = true

func suspendProcess() { _ = "STUB: not implemented"; return }
