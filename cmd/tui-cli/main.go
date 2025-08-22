package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed templates/button.tmpl
var buttonTmpl string

func usage() {
	fmt.Println("tui-cli init <appdir>\n" +
		"tui-cli add button <destdir>")
}

func ensureDir(dir string) error {
	return os.MkdirAll(dir, 0o755)
}

func writeFile(path, content string) error {
	if err := ensureDir(filepath.Dir(path)); err != nil { return err }
	return os.WriteFile(path, []byte(content), 0o644)
}

func cmdInit(args []string) error {
	if len(args) < 1 { return fmt.Errorf("missing <appdir>") }
	app := args[0]
	// Write theme skeleton into app
	tokens := `package theme

type Tokens struct {
	Colors struct {
		Bg, Surface, Text, Muted, Primary, PrimaryFg, Warning, Success string
	}
	Space  []int
	Radius []int
	Border struct{ Normal, Focused string }
}

var LightTokens = func() Tokens {
	var t Tokens
	t.Colors = struct {
		Bg, Surface, Text, Muted, Primary, PrimaryFg, Warning, Success string
	}{"#0B0C0F", "#111317", "#E6E8EB", "#9BA3AF", "#3B82F6", "#0B0C0F", "#F59E0B", "#10B981"}
	t.Space = []int{0,1,2,3,4,6,8}
	t.Radius = []int{0,1,2,3}
	t.Border = struct{ Normal, Focused string }{"#2A2F3A", "#3B82F6"}
	return t
}()
`
	styles := `package theme

import "github.com/charmbracelet/lipgloss"

type Styles struct {
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
	return s
}
`
	themeGo := `package theme

type Theme struct {
	Tokens Tokens
	Styles Styles
	Name   string
}

func New(name string, tokens Tokens) Theme {
	return Theme{Name: name, Tokens: tokens, Styles: BuildStyles(tokens)}
}
`
	if err := writeFile(filepath.Join(app, "theme", "tokens.go"), tokens); err != nil { return err }
	if err := writeFile(filepath.Join(app, "theme", "styles.go"), styles); err != nil { return err }
	if err := writeFile(filepath.Join(app, "theme", "theme.go"), themeGo); err != nil { return err }
	fmt.Println("initialized theme in:", filepath.Join(app, "theme"))
	return nil
}

func cmdAdd(args []string) error {
	if len(args) < 2 { return fmt.Errorf("usage: tui-cli add button <destdir>") }
	comp := args[0]
	dest := args[1]
	switch comp {
	case "button":
		path := filepath.Join(dest, "button.go")
		if err := writeFile(path, buttonTmpl); err != nil { return err }
		fmt.Println("wrote:", path)
	default:
		return fmt.Errorf("unknown component: %s", comp)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 { usage(); return }
	switch os.Args[1] {
	case "init":
		if err := cmdInit(os.Args[2:]); err != nil { fmt.Println("error:", err); os.Exit(1) }
	case "add":
		if err := cmdAdd(os.Args[2:]); err != nil { fmt.Println("error:", err); os.Exit(1) }
	default:
		usage()
	}
}
