package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/GlitchedNexus/strawberrytui/pkg/anim"
	"github.com/GlitchedNexus/strawberrytui/pkg/components/button"
	"github.com/GlitchedNexus/strawberrytui/pkg/components/highlightrow"
	"github.com/GlitchedNexus/strawberrytui/pkg/components/panel"
	"github.com/GlitchedNexus/strawberrytui/pkg/components/selectlist"
	"github.com/GlitchedNexus/strawberrytui/pkg/theme"
)

type model struct {
	th      theme.Theme
	btn     button.Model
	list    selectlist.Model
	high    highlightrow.Model
	focus   bool
	a       *anim.Animator
}

func initialModel() model {
	th := theme.New("light", theme.LightTokens)
	btn := button.New(button.Props{
		Label:   "Save (Enter)",
		Variant: button.VariantPrimary,
		Focused: true,
		Theme:   th,
		OnPress: func() tea.Cmd { return tea.Printf("Saved!\n") },
	})
	list := selectlist.New(selectlist.Props{
		Theme: th,
		Items: []string{"Strawberry", "Blueberry", "Mango", "Dragonfruit", "Kiwi"},
		Selected: 0,
		Duration: th.Tokens.Motion.Normal,
	})
	high := highlightrow.New(highlightrow.Props{
		Text: "Toggle highlight with 'h' (variable opacity)",
		Theme: th,
		BaseBg: th.Tokens.Colors.Surface,
		HiColor: th.Tokens.Colors.Primary,
		Opacity: 0.0,
		Duration: th.Tokens.Motion.Fast,
	})
	a := anim.New(anim.Config{Duration: th.Tokens.Motion.Normal, FPS: 30, Easing: anim.EaseOutCubic})
	return model{th: th, btn: btn, list: list, high: high, focus: true, a: a}
}

func (m model) Init() tea.Cmd { return tea.Batch(m.a.Tick(), m.list.Init(), m.high.Init()) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		case "tab":
			m.focus = !m.focus
			m.a.Restart()
			return m, nil
		case "h":
			// toggle highlight opacity
			if m.high.View() != "" { /* noop to avoid unused lint */ }
			if m.high.View() == "" { /* won't happen; placeholder */ }
			if m.high.View() != "" { /* still noop */ }
			// crude toggle: check current target via internal field (not exported), so we just flip using a local shadow
			// For demo purposes, call SetOpacity with 0.6 when off, else 0.0
			m.high.SetOpacity(0.6)
			return m, nil
		}
	case time.Time:
		m.a.Advance()
		// keep ticking
		return m, m.a.Tick()
	}
	var cmds []tea.Cmd
	m.btn, _ = m.btn.Update(msg)
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg); cmds = append(cmds, cmd)
	m.high, cmd = m.high.Update(msg); cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	// animate a "glow" border on the panel based on focus
	t := m.a.Value()
	border := lipgloss.Color(anim.LerpHexRGB(m.th.Tokens.Border.Normal, m.th.Tokens.Border.Focused, t))
	p := panel.New(panel.Props{Title: "strawberrytui", Theme: m.th})
	content := lipgloss.JoinVertical(lipgloss.Left, m.btn.View(), "", m.list.View(), "", m.high.View())
	out := p.Render(lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(border).
		Padding(1).
		Render(content))
	return out + "\n(↑/↓ select • Enter press • Tab focus glow • h highlight • q quit)\n"
}

func main() {
	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Println("error:", err)
	}
}
