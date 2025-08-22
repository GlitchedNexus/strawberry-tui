package selectlist

import (
	"strings"
	"time"

	"github.com/GlitchedNexus/strawberrytui/pkg/theme"
	"github.com/GlitchedNexus/strawberrytui/pkg/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Props struct {
	Theme theme.Theme
	Items []string
	Selected int
	Duration time.Duration // animation duration
}

type Model struct {
	p Props
	keymap struct{ up, down, selectK key.Binding }
	animT float64 // 0..1 progress for current selection
	animDir float64
	lastTick time.Time
}

func New(p Props) Model {
	if p.Duration == 0 { p.Duration = 200 * time.Millisecond }
	m := Model{ p: p, animT: 1.0 }
	m.keymap.up = key.NewBinding(key.WithKeys("up", "k"))
	m.keymap.down = key.NewBinding(key.WithKeys("down", "j"))
	m.keymap.selectK = key.NewBinding(key.WithKeys("enter"))
	return m
}

func (m Model) Init() tea.Cmd { return tick() }

func tick() tea.Cmd { return tea.Tick(1*time.Second/30, func(t time.Time) tea.Msg { return t }) }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keymap.up) {
			if m.p.Selected > 0 { m.p.Selected--; m.animDir = -1; m.animT = 0 }
			return m, nil
		}
		if key.Matches(msg, m.keymap.down) {
			if m.p.Selected < len(m.p.Items)-1 { m.p.Selected++; m.animDir = +1; m.animT = 0 }
			return m, nil
		}
	case time.Time:
		step := 1.0 / (float64(m.p.Duration) / (1000.0/30.0))
		if m.animT < 1.0 { m.animT += step; if m.animT > 1 { m.animT = 1 } }
		return m, tick()
	}
	return m, nil
}

func easeOutCubic(t float64) float64 { u := 1 - t; return 1 - u*u*u }

func (m Model) View() string {
	var b strings.Builder
	padBase := 1
	padMax := 4
	styles := m.p.Theme.Styles
	for i, it := range m.p.Items {
		selected := i == m.p.Selected
		s := styles.Item
		if selected {
			// animate padding & color
			f := easeOutCubic(m.animT)
			pad := padBase + int(f*float64(padMax-padBase))
			bg := utils.LerpHex(m.p.Theme.Tokens.Colors.Surface, m.p.Theme.Tokens.Colors.Primary, f)
			fg := utils.LerpHex(m.p.Theme.Tokens.Colors.Text,    m.p.Theme.Tokens.Colors.PrimaryFg, f)
			s = lipgloss.NewStyle().
				Padding(0, pad).
				Background(lipgloss.Color(bg)).
				Foreground(lipgloss.Color(fg)).
				Bold(true)
		}
		b.WriteString(s.Render(it))
		b.WriteString("\n")
	}
	return b.String()
}
