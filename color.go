package tea

import (
	"image/color"
)

type backgroundColorMsg struct{}

func RequestBackgroundColor() Msg { _ = "STUB: not implemented"; return *new(Msg) }

type foregroundColorMsg struct{}

func RequestForegroundColor() Msg { _ = "STUB: not implemented"; return *new(Msg) }

type cursorColorMsg struct{}

func RequestCursorColor() Msg { _ = "STUB: not implemented"; return *new(Msg) }

type ForegroundColorMsg struct{ color.Color }

func (e ForegroundColorMsg) String() string { _ = "STUB: not implemented"; return "" }

func (e ForegroundColorMsg) IsDark() bool { _ = "STUB: not implemented"; return false }

type BackgroundColorMsg struct{ color.Color }

func (e BackgroundColorMsg) String() string { _ = "STUB: not implemented"; return "" }

func (e BackgroundColorMsg) IsDark() bool { _ = "STUB: not implemented"; return false }

type CursorColorMsg struct{ color.Color }

func (e CursorColorMsg) String() string { _ = "STUB: not implemented"; return "" }

func (e CursorColorMsg) IsDark() bool { _ = "STUB: not implemented"; return false }
