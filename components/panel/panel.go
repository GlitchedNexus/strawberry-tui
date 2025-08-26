package panel

import (
	"strings"
	"github.com/charmbracelet/lipgloss"
	"github.com/GlitchedNexus/strawberrytui/pkg/theme"
)

type Props struct {
	Title string
	Theme theme.Theme
}

type Model struct {
	p      Props
	styles struct{ base, header lipgloss.Style }
}

func New(p Props) Model {
	m := Model{p: p}
	m.styles.base = p.Theme.Styles.Panel.Base
	m.styles.header = p.Theme.Styles.Panel.Header
	return m
}

func (m Model) Render(content string) string {
	head := m.styles.header.Render(" " + m.p.Title + " ")
	body := m.styles.base.Render(content)
	return strings.Join([]string{head, body}, "\n")
}
