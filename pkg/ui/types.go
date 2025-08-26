package ui

// Color is a terminal color index (-1 = default).
type Color int16

// Attr are per-cell text attributes (kept minimal on purpose).
type Attr struct {
	FG, BG   Color
	Bold     bool
	Underline bool
}

// Rect is a cell-space rectangle.
type Rect struct{ X, Y, W, H int }
