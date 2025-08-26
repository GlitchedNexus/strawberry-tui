package selectlist

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
	"github.com/GlitchedNexus/strawberrytui/pkg/anim"
	"github.com/GlitchedNexus/strawberrytui/pkg/theme"
)

// Props configure the SelectList.
type Props struct {
	Theme    theme.Theme
	Items    []string
	Selected int
	Duration time.Duration // optional; defaults to Theme.Tokens.Motion.Normal
}

// Model renders an animated keyboard-driven list.
type Model struct {
	p Props
	a *anim.Animator
	keymap struct{ up, down, sel key.Binding }
}

func New(p Props) Model {
	if p.Duration == 0 { p.Duration = p.Theme.Tokens.Motion.Normal }
	m := Model{ p: p, a: anim.New(anim.Config{Duration: p.Duration, FPS: 30, Easing: anim.EaseOutCubic}) }
	m.keymap.up = key.NewBinding(key.WithKeys("up", "k"))
	m.keymap.down = key.NewBinding(key.WithKeys("down", "j"))
	m.keymap.sel = key.NewBinding(key.WithKeys("enter"))
	return m
}

func (m Model) Init() tea.Cmd { return m.a.Tick() }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keymap.up) {
			if m.p.Selected > 0 { m.p.Selected--; m.a.Restart() }
			return m, nil
		}
		if key.Matches(msg, m.keymap.down) {
			if m.p.Selected < len(m.p.Items)-1 { m.p.Selected++; m.a.Restart() }
			return m, nil
		}
	case time.Time:
		m.a.Advance()
		return m, m.a.Tick()
	}
	return m, nil
}

func (m Model) View() string {
	var b strings.Builder
	t := m.a.Value()
	padBase, padMax := 1, 4
	for i, it := range m.p.Items {
		selected := i == m.p.Selected
		s := m.p.Theme.Styles.Highlighter // will override for unselected
		if selected {
			pad := anim.LerpInt(padBase, padMax, t)
			bg := anim.LerpHexRGB(m.p.Theme.Tokens.Colors.Surface, m.p.Theme.Tokens.Colors.Primary, t)
			fg := anim.LerpHexRGB(m.p.Theme.Tokens.Colors.Text,    m.p.Theme.Tokens.Colors.PrimaryFg, t)
			s = lipgloss.NewStyle().
				Padding(0, pad).
				Background(lipgloss.Color(bg)).
				Foreground(lipgloss.Color(fg)).
				Bold(true)
		} else {
			s = m.p.Theme.Styles.Panel.Base.Padding(0,1)
		}
		b.WriteString(s.Render(it))
		b.WriteString("\n")
	}
	return b.String()
}
