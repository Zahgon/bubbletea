package tea

import (
	"time"
)

func Batch(cmds ...Cmd) Cmd { _ = "STUB: not implemented"; return *new(Cmd) }

type BatchMsg []Cmd

func Sequence(cmds ...Cmd) Cmd { _ = "STUB: not implemented"; return *new(Cmd) }

type sequenceMsg []Cmd

func compactCmds[T ~[]Cmd](cmds []Cmd) Cmd { _ = "STUB: not implemented"; return *new(Cmd) }

func Every(duration time.Duration, fn func(time.Time) Msg) Cmd {
	_ = "STUB: not implemented"
	return *new(Cmd)
}

func Tick(d time.Duration, fn func(time.Time) Msg) Cmd { _ = "STUB: not implemented"; return *new(Cmd) }

type windowSizeMsg struct{}

func RequestWindowSize() Msg { _ = "STUB: not implemented"; return *new(Msg) }
