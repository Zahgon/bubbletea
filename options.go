package tea

import (
	"context"
	"io"

	"github.com/charmbracelet/colorprofile"
)

type ProgramOption func(*Program)

func WithContext(ctx context.Context) ProgramOption {
	_ = "STUB: not implemented"
	return *new(ProgramOption)
}

func WithOutput(output io.Writer) ProgramOption {
	_ = "STUB: not implemented"
	return *new(ProgramOption)
}

func WithInput(input io.Reader) ProgramOption {
	_ = "STUB: not implemented"
	return *new(ProgramOption)
}

func WithEnvironment(env []string) ProgramOption {
	_ = "STUB: not implemented"
	return *new(ProgramOption)
}

func WithoutSignalHandler() ProgramOption { _ = "STUB: not implemented"; return *new(ProgramOption) }

func WithoutCatchPanics() ProgramOption { _ = "STUB: not implemented"; return *new(ProgramOption) }

func WithoutSignals() ProgramOption { _ = "STUB: not implemented"; return *new(ProgramOption) }

func WithoutRenderer() ProgramOption { _ = "STUB: not implemented"; return *new(ProgramOption) }

func WithFilter(filter func(Model, Msg) Msg) ProgramOption {
	_ = "STUB: not implemented"
	return *new(ProgramOption)
}

func WithFPS(fps int) ProgramOption { _ = "STUB: not implemented"; return *new(ProgramOption) }

func WithColorProfile(profile colorprofile.Profile) ProgramOption {
	_ = "STUB: not implemented"
	return *new(ProgramOption)
}

func WithWindowSize(width, height int) ProgramOption {
	_ = "STUB: not implemented"
	return *new(ProgramOption)
}
