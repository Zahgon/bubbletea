//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || aix || zos
// +build darwin dragonfly freebsd linux netbsd openbsd solaris aix zos

package tea

func (p *Program) listenForResize(done chan struct{}) { _ = "STUB: not implemented"; return }
