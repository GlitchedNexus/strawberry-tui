package theme

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Item   lipgloss.Style
	ItemSelected lipgloss.Style
	Accent lipgloss.Style
	Button struct {
		Base, Primary, Ghost lipgloss.Style
		Focused              lipgloss.Style
	}
	Panel struct {
		Base, Header lipgloss.Style
	}
}

func BuildStyles(t Tokens) Styles {
	var s Styles
	text := lipgloss.Color(t.Colors.Text)
	surf := lipgloss.Color(t.Colors.Surface)
	primary := lipgloss.Color(t.Colors.Primary)
	accent := lipgloss.Color(t.Border.Focused)
	bg := lipgloss.Color(t.Colors.Bg)
	primaryFg := lipgloss.Color(t.Colors.PrimaryFg)
	border := lipgloss.Color(t.Border.Normal)
	focus := lipgloss.Color(t.Border.Focused)

	s.Item = lipgloss.NewStyle().Foreground(text).Background(surf).Padding(0,1)
	s.ItemSelected = lipgloss.NewStyle().Foreground(lipgloss.Color(t.Colors.PrimaryFg)).Background(primary).Padding(0,2)
	s.Accent = lipgloss.NewStyle().Foreground(lipgloss.Color(accent))


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
	return s
}
