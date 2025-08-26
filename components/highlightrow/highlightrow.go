package highlightrow

import (
	"time"

	"github.com/GlitchedNexus/strawberry-tui/pkg/anim"
	"github.com/GlitchedNexus/strawberry-tui/pkg/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Props configure the row highlighting with variable opacity.
type Props struct {
	Text     string
	Theme    theme.Theme
	BaseBg   string // usually th.Tokens.Colors.Surface
	HiColor  string // usually th.Tokens.Colors.Primary
	Opacity  float64 // target 0..1
	Duration time.Duration // defaults to Theme.Tokens.Motion.Normal
}

type Model struct {
	p Props
	a *anim.Animator
}

func New(p Props) Model {
	if p.Duration == 0 { p.Duration = p.Theme.Tokens.Motion.Normal }
	m := Model{ p: p, a: anim.New(anim.Config{Duration: p.Duration, FPS: 30, Easing: anim.EaseOutCubic}) }
	return m
}

func (m *Model) SetOpacity(alpha float64) { m.p.Opacity = alpha; m.a.Restart() }

func (m Model) Init() tea.Cmd { return m.a.Tick() }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg.(type) {
	case time.Time:
		m.a.Advance()
		return m, m.a.Tick()
	}
	return m, nil
}

func (m Model) View() string {
	// ease toward target opacity
	t := m.a.Value()
	alpha := t * m.p.Opacity
	bg := lipgloss.Color(anim.LerpHexRGB(m.p.BaseBg, m.p.HiColor, alpha))
	st := lipgloss.NewStyle().
		Background(bg).
		Foreground(lipgloss.Color(m.p.Theme.Tokens.Colors.Text)).
		Padding(0, 1)
	return st.Render(m.p.Text)
}
