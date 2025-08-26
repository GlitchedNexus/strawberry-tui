package theme

import "github.com/charmbracelet/lipgloss"

// Theme bundles Tokens and exposes resolvers for both render paths.
type Theme struct {
	Name   string
	Tokens Tokens
}

func Default() Theme { return Theme{Name: "strawberry", Tokens: DefaultTokens()} }

// ---------------- Immediate-mode resolver (Lipgloss) ----------------

func (th Theme) ResolveLipgloss(base lipgloss.Style, spec StyleSpec) lipgloss.Style {
	s := base
	// Colors (token or #hex)
	if spec.FGHex   != "" { s = s.Foreground(lipgloss.Color(spec.FGHex)) }
	if spec.FGToken != "" { s = s.Foreground(lipgloss.Color(th.Tokens.Color(spec.FGToken))) }
	if spec.BGHex   != "" { s = s.Background(lipgloss.Color(spec.BGHex)) }
	if spec.BGToken != "" { s = s.Background(lipgloss.Color(th.Tokens.Color(spec.BGToken))) }
	if spec.BorderHex   != "" { s = s.BorderForeground(lipgloss.Color(spec.BorderHex)) }
	if spec.BorderToken != "" { s = s.BorderForeground(lipgloss.Color(th.Tokens.Color(spec.BorderToken))) }

	// Padding
	pt, pr, pb, pl := padFrom(spec)
	s = s.Padding(pt, pr, pb, pl)

	// Text attrs
	if spec.Bold != nil      { s = s.Bold(*spec.Bold) }
	if spec.Underline != nil { s = s.Underline(*spec.Underline) }

	// Note: Lipgloss doesn't have corner radius; choose border style per radius if desired.
	return s
}

func padFrom(spec StyleSpec) (pt, pr, pb, pl int) {
	if spec.P  != nil { pt, pr, pb, pl = *spec.P, *spec.P, *spec.P, *spec.P }
	if spec.Px != nil { pr, pl = *spec.Px, *spec.Px }
	if spec.Py != nil { pt, pb = *spec.Py, *spec.Py }
	if spec.Pt != nil { pt = *spec.Pt }
	if spec.Pr != nil { pr = *spec.Pr }
	if spec.Pb != nil { pb = *spec.Pb }
	if spec.Pl != nil { pl = *spec.Pl }
	return
}

// ---------------- Retained-mode resolver (to your UI attrs) ----------------

// Adapt these to your retained-mode types.
type ResolvedTUI struct {
	Attr struct {
		FG, BG int16
		Bold, Underline bool
	}
	Padding struct{ T, R, B, L int }
	Radius    int
	BorderHex string
}

func (th Theme) ResolveTUI(spec StyleSpec) ResolvedTUI {
	var r ResolvedTUI

	// Resolve colors: prefer explicit hex, else token lookup
	fg := pickHex(spec.FGHex,   th.Tokens.Color(spec.FGToken))
	bg := pickHex(spec.BGHex,   th.Tokens.Color(spec.BGToken))
	bc := pickHex(spec.BorderHex, th.Tokens.Color(spec.BorderToken))

	r.Attr.FG = toTermColor(fg) // TODO: hex→256-color mapping (or -1 for default)
	r.Attr.BG = toTermColor(bg)
	r.BorderHex = bc

	if spec.Bold != nil      { r.Attr.Bold      = *spec.Bold }
	if spec.Underline != nil { r.Attr.Underline = *spec.Underline }

	pt, pr, pb, pl := padFrom(spec)
	r.Padding = struct{T,R,B,L int}{pt, pr, pb, pl}

	// Radius from class or token
	if spec.Radius != nil { r.Radius = *spec.Radius
	} else if spec.RadiusKey != "" { r.Radius = th.Tokens.RadiusVal(spec.RadiusKey)
	}

	return r
}

func pickHex(hex, tokenHex string) string { if hex != "" { return hex }; return tokenHex }
func toTermColor(hex string) int16        { if hex == "" { return -1 }; return -1 /* map later */ }
