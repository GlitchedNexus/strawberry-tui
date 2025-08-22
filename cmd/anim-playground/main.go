package main

import (
	"fmt"

	"github.com/GlitchedNexus/strawberrytui/pkg/components/selectlist"
	"github.com/GlitchedNexus/strawberrytui/pkg/theme"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list selectlist.Model
}

func initial() model {
	th := theme.New("light", theme.LightTokens)
	l := selectlist.New(selectlist.Props{
		Theme: th,
		Items: []string{"Strawberry", "Blueberry", "Mango", "Dragonfruit", "Kiwi"},
		Selected: 0,
	})
	return model{list: l}
}

func (m model) Init() tea.Cmd { return m.list.Init() }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View() + "\n↑/↓ select • Enter choose • q to quit\n"
}

func main() {
	p := tea.NewProgram(initial())
	if _, err := p.Run(); err != nil { fmt.Println("error:", err) }
}
