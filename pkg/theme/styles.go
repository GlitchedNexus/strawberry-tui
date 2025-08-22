package theme

import "github.com/charmbracelet/lipgloss"

// Styles groups lipgloss styles derived from Tokens.
type Styles struct {
	Button struct {
		Base, Primary, Ghost lipgloss.Style
		Focused              lipgloss.Style
	}
	Panel struct {
		Base, Header lipgloss.Style
	}
	Highlighter lipgloss.Style
}

// BuildStyles constructs Styles from Tokens.
func BuildStyles(t Tokens) Styles {
	var s Styles
	bg := lipgloss.Color(t.Colors.Bg)
	surf := lipgloss.Color(t.Colors.Surface)
	text := lipgloss.Color(t.Colors.Text)
	primary := lipgloss.Color(t.Colors.Primary)
	primaryFg := lipgloss.Color(t.Colors.PrimaryFg)
	border := lipgloss.Color(t.Border.Normal)
	focus := lipgloss.Color(t.Border.Focused)

	s.Button.Base = lipgloss.NewStyle().
		Padding(0, 2).
		Foreground(text).
		Background(surf).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(border)

	s.Button.Primary = s.Button.Base.
		Background(primary).
		Foreground(primaryFg)

	s.Button.Ghost = lipgloss.NewStyle().
		Padding(0, 2).
		Foreground(text)

	s.Button.Focused = s.Button.Base.BorderForeground(focus)

	s.Panel.Base = lipgloss.NewStyle().
		Background(surf).
		Foreground(text).
		Padding(t.Space[2])
	s.Panel.Header = lipgloss.NewStyle().Bold(true).Foreground(bg)

	// Default highlighter (used by HighlightRow component)
	s.Highlighter = lipgloss.NewStyle().
		Background(primary).
		Foreground(primaryFg)

	return s
}
