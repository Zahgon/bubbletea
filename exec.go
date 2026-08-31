package tea

import (
	"io"
	"os/exec"
)

type execMsg struct {
	cmd ExecCommand
	fn  ExecCallback
}

func Exec(c ExecCommand, fn ExecCallback) Cmd { _ = "STUB: not implemented"; return *new(Cmd) }

func ExecProcess(c *exec.Cmd, fn ExecCallback) Cmd { _ = "STUB: not implemented"; return *new(Cmd) }

type ExecCallback func(error) Msg

type ExecCommand interface {
	Run() error
	SetStdin(io.Reader)
	SetStdout(io.Writer)
	SetStderr(io.Writer)
}

func wrapExecCommand(c *exec.Cmd) ExecCommand { _ = "STUB: not implemented"; return *new(ExecCommand) }

type osExecCommand struct{ *exec.Cmd }

func (c *osExecCommand) SetStdin(r io.Reader) { _ = "STUB: not implemented"; return }

func (c *osExecCommand) SetStdout(w io.Writer) { _ = "STUB: not implemented"; return }

func (c *osExecCommand) SetStderr(w io.Writer) { _ = "STUB: not implemented"; return }

func (p *Program) exec(c ExecCommand, fn ExecCallback) { _ = "STUB: not implemented"; return }
