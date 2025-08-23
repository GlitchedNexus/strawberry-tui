package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed templates/button.tmpl
var buttonTmpl string

//go:embed templates/selectlist.tmpl
var selectListTmpl string

func usage() {
	fmt.Println("tui-cli init <appdir>\n" +
		"tui-cli add button <destdir>\n" +
		"tui-cli add selectlist <destdir>")
}

func ensureDir(dir string) error { return os.MkdirAll(dir, 0o755) }
func writeFile(path, content string) error {
	if err := ensureDir(filepath.Dir(path)); err != nil { return err }
	return os.WriteFile(path, []byte(content), 0o644)
}

func cmdInit(args []string) error {
	if len(args) < 1 { return fmt.Errorf("missing <appdir>") }
	app := args[0]
	tokens := `package theme

type Tokens struct {
	Colors struct {
		Bg, Surface, Text, Muted, Primary, PrimaryFg, Warning, Success string
	}
	Space  []int
	Radius []int
	Border struct{ Normal, Focused string }
	Motion Motion
}

type Motion struct { Fast, Normal, Slow time.Duration }
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
	Highlighter lipgloss.Style
}
`
	themeGo := `package theme
type Theme struct { Tokens Tokens; Styles Styles; Name string }
func New(name string, tokens Tokens) Theme { return Theme{Name: name, Tokens: tokens, Styles: BuildStyles(tokens)} }
`
	if err := writeFile(filepath.Join(app, "theme", "tokens.go"), tokens); err != nil { return err }
	if err := writeFile(filepath.Join(app, "theme", "styles.go"), styles); err != nil { return err }
	if err := writeFile(filepath.Join(app, "theme", "theme.go"), themeGo); err != nil { return err }
	fmt.Println("initialized theme in:", filepath.Join(app, "theme"))
	return nil
}

func cmdAdd(args []string) error {
	if len(args) < 2 { return fmt.Errorf("usage: tui-cli add <component> <destdir>") }
	comp := args[0]; dest := args[1]
	switch comp {
		case "button":
			return writeFile(filepath.Join(dest, "button.go"), buttonTmpl)
		case "selectlist":
			return writeFile(filepath.Join(dest, "selectlist.go"), selectListTmpl)
		default:
			return fmt.Errorf("unknown component: %s", comp)
	}
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
