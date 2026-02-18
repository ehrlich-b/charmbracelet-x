package vt

import (
	"image/color"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
)

// Callbacks represents a set of callbacks for a terminal.
type Callbacks struct {
	// Bell callback. When set, this function is called when a bell character is
	// received.
	Bell func()

	// Title callback. When set, this function is called when the terminal title
	// changes.
	Title func(string)

	// IconName callback. When set, this function is called when the terminal
	// icon name changes.
	IconName func(string)

	// AltScreen callback. When set, this function is called when the alternate
	// screen is activated or deactivated.
	AltScreen func(bool)

	// CursorPosition callback. When set, this function is called when the cursor
	// position changes.
	CursorPosition func(old, new uv.Position) //nolint:predeclared,revive

	// CursorVisibility callback. When set, this function is called when the
	// cursor visibility changes.
	CursorVisibility func(visible bool)

	// CursorStyle callback. When set, this function is called when the cursor
	// style changes.
	CursorStyle func(style CursorStyle, blink bool)

	// CursorColor callback. When set, this function is called when the cursor
	// color changes. Nil indicates the default terminal color.
	CursorColor func(color color.Color)

	// BackgroundColor callback. When set, this function is called when the
	// background color changes. Nil indicates the default terminal color.
	BackgroundColor func(color color.Color)

	// ForegroundColor callback. When set, this function is called when the
	// foreground color changes. Nil indicates the default terminal color.
	ForegroundColor func(color color.Color)

	// WorkingDirectory callback. When set, this function is called when the
	// current working directory changes.
	WorkingDirectory func(string)

	// EnableMode callback. When set, this function is called when a mode is
	// enabled.
	EnableMode func(mode ansi.Mode)

	// DisableMode callback. When set, this function is called when a mode is
	// disabled.
	DisableMode func(mode ansi.Mode)

	// ScrollOut is called when lines scroll off the top of the scroll region.
	// Provides the line data BEFORE destruction. Only fires when the cursor is
	// at the top of the scroll region (i.e. during ScrollUp and CSI S, but not
	// mid-screen CSI M). The caller can render these lines (via Line.Render())
	// for scrollback storage. Callers that only want main-screen scrollback
	// should check [Emulator.IsAltScreen] and ignore while in alt screen.
	ScrollOut func(lines []uv.Line)

	// ScrollbackClear is called when the terminal requests scrollback erasure.
	// Fired on ESC[3J (Erase Scrollback) and ESC c (Full Reset / RIS).
	ScrollbackClear func()
}
