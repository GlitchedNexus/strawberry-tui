package button

import (
	"github.com/GlitchedNexus/strawberry-tui/pkg/theme"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Variant int
const (
	VariantBase Variant = iota
	VariantPrimary
	VariantGhost
)

type Props struct {
	Label   string
	Variant Variant
	Focused bool
	OnPress func() tea.Cmd
	Theme   theme.Theme
}

type Model struct {
	p      Props
	styles struct{ base, focused lipgloss.Style }
	keymap struct{ submit key.Binding }
}

func New(p Props) Model {
	m := Model{p: p}
	switch p.Variant {
	case VariantPrimary:
		m.styles.base = p.Theme.Styles.Button.Primary
	case VariantGhost:
		m.styles.base = p.Theme.Styles.Button.Ghost
	default:
		m.styles.base = p.Theme.Styles.Button.Base
	}
	m.styles.focused = p.Theme.Styles.Button.Focused
	m.keymap.submit = key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "activate"))
	return m
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keymap.submit) && m.p.Focused && m.p.OnPress != nil {
			return m, m.p.OnPress()
		}
	}
	return m, nil
}

func (m Model) View() string {
	st := m.styles.base
	if m.p.Focused { st = m.styles.focused }
	return st.Render(m.p.Label)
}
