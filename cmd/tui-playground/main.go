package main

import (
	"fmt"

	"github.com/GlitchedNexus/strawberrytui/pkg/components/button"
	"github.com/GlitchedNexus/strawberrytui/pkg/components/panel"
	"github.com/GlitchedNexus/strawberrytui/pkg/theme"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	th   theme.Theme
	btn  button.Model
	quit bool
}

func initialModel() model {
	th := theme.New("light", theme.LightTokens)
	btn := button.New(button.Props{
		Label:   "Save",
		Variant: button.VariantPrimary,
		Focused: true,
		Theme:   th,
		OnPress: func() tea.Cmd { return tea.Printf("Saved!\n") },
	})
	return model{th: th, btn: btn}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			m.quit = true
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.btn, cmd = m.btn.Update(msg)
	return m, cmd
}

func (m model) View() string {
	p := panel.New(panel.Props{Title: "strawberrytui", Theme: m.th})
	return p.Render(m.btn.View()) + "\n\n" + "(q to quit)\n"
}

func main() {
	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Println("error:", err)
	}
}
