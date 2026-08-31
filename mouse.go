package tea

import (
	"fmt"

	uv "github.com/charmbracelet/ultraviolet"
)

type MouseButton = uv.MouseButton

const (
	MouseNone       = uv.MouseNone
	MouseLeft       = uv.MouseLeft
	MouseMiddle     = uv.MouseMiddle
	MouseRight      = uv.MouseRight
	MouseWheelUp    = uv.MouseWheelUp
	MouseWheelDown  = uv.MouseWheelDown
	MouseWheelLeft  = uv.MouseWheelLeft
	MouseWheelRight = uv.MouseWheelRight
	MouseBackward   = uv.MouseBackward
	MouseForward    = uv.MouseForward
	MouseButton10   = uv.MouseButton10
	MouseButton11   = uv.MouseButton11
)

type MouseMsg interface {
	fmt.Stringer

	Mouse() Mouse
}

type Mouse struct {
	X, Y   int
	Button MouseButton
	Mod    KeyMod
}

func (m Mouse) String() (s string) { _ = "STUB: not implemented"; return "" }

type MouseClickMsg Mouse

func (e MouseClickMsg) String() string { _ = "STUB: not implemented"; return "" }

func (e MouseClickMsg) Mouse() Mouse { _ = "STUB: not implemented"; return *new(Mouse) }

type MouseReleaseMsg Mouse

func (e MouseReleaseMsg) String() string { _ = "STUB: not implemented"; return "" }

func (e MouseReleaseMsg) Mouse() Mouse { _ = "STUB: not implemented"; return *new(Mouse) }

type MouseWheelMsg Mouse

func (e MouseWheelMsg) String() string { _ = "STUB: not implemented"; return "" }

func (e MouseWheelMsg) Mouse() Mouse { _ = "STUB: not implemented"; return *new(Mouse) }

type MouseMotionMsg Mouse

func (e MouseMotionMsg) String() string { _ = "STUB: not implemented"; return "" }

func (e MouseMotionMsg) Mouse() Mouse { _ = "STUB: not implemented"; return *new(Mouse) }
