# **StrawberryUI**

A cute component library for you TUI applications built with BubbleTea :D

---

## 📦 Repository Layout

```
strawberryui/
├── cmd/
│   └── scaffold/               # CLI tool to generate new components or themes
├── pkg/
│   ├── core/                   # Core types and interfaces
│   │   ├── component.go        # Component interface definitions
│   │   ├── theme.go            # Theme interface and default palette
│   │   └── style.go            # Style tokens & helpers (spacing, borders, colors)
│   ├── components/             # Built-in components
│   │   ├── button/             # Button component
│   │   ├── input/              # TextInput component
│   │   ├── list/               # List & ListItem components
│   │   ├── table/              # Table component
│   │   └── dialog/             # Modal/dialog components
│   ├── themes/                 # Predefined themes (light, dark, high-contrast)
│   │   ├── light.go
│   │   ├── dark.go
│   │   └── contrast.go
│   └── utils/                  # Utilities (alignment, truncation, responsive layouts)
│       ├── layout.go
│       └── helpers.go
├── examples/                   # Sample Bubble Tea apps demonstrating StrawberryUI
│   ├── counter/
│   └── todo/
├── docs/                       # Markdown docs & component usage guides
│   ├── getting-started.md
│   ├── theming.md
│   └── api-reference.md
├── .github/                    # CI workflows, issue templates
│   └── workflows/
│       └── ci.yml
├── go.mod
└── README.md
```

---

## 🎯 Design Goals

1. **Composable**
   Each UI element implements a small `Component` interface so you can nest, wrap, or combine them freely in any Bubble Tea model.

2. **Themeable**
   Centralized tokens for colors, borders, spacing; swappable themes at runtime.

3. **Minimal Overhead**
   Depend only on Bubble Tea, Bubbles and Lip Gloss—no extra heavy frameworks.

4. **Discoverable & Documented**
   Well-structured docs and examples; a CLI scaffold to jump-start new components or custom themes.

---

## 🏗️ Core Interfaces

```go
// pkg/core/component.go
package core

type Component interface {
    // Render returns the styled string for this component.
    Render() string
    // SetSize lets layout logic inform the component of available width/height.
    SetSize(width, height int)
}
```

```go
// pkg/core/theme.go
package core

type Theme interface {
    PrimaryColor() string
    SecondaryColor() string
    BorderStyle() string
    Spacing(unit int) int
    // …etc.
}
```

```go
// pkg/core/style.go
package core

// StyleToken holds raw lipgloss.Style for reuse and composition.
type StyleToken struct {
    Base  lipgloss.Style
    Theme Theme
}

func (s *StyleToken) WithForeground(c string) lipgloss.Style { … }
func (s *StyleToken) WithPadding(u int) lipgloss.Style         { … }
```

---

## 🧩 Built-in Components

### Button

* **API**

  ```go
  type Button struct { /* label, style tokens, focus state */ }
  func NewButton(label string, opts ...ButtonOption) *Button
  func (b *Button) Render() string
  ```
* **Options**

  * `WithStyle(lipgloss.Style)`
  * `OnPress(msg tea.Msg)`

### TextInput

* Wraps `bubbles.TextInput` but applies StrawberryUI styles.
* Exposes options: placeholder, width, cursor style.

### List & ListItem

* Horizontal or vertical lists with selectable items.
* Supports pagination, keybindings, custom item renderers.

### Table

* Grid layout with headers, column alignment, optional cell padding.
* Features: zebra striping, scrollable overflow.

### Dialog / Modal

* Centered overlay with backdrop, title, body, footer buttons.
* Animations via `lipgloss.Place` helpers.

---

## 🎨 Theming

* **Default Themes**

  * `themes/light.go` and `themes/dark.go` provide `var Default = Theme{…}`
* **Custom**

  ```go
  myTheme := theme.New(
    theme.WithPrimary("#FF3366"),
    theme.WithBorderStyle("rounded"),
    theme.WithSpacing(1),
  )
  ```
* **Runtime switching**
  Store active theme in `core.CurrentTheme` and re-render components on change.

---

## 🔄 Rendering Pipeline

1. **Model** holds component tree.
2. On `View()`, call `rootComponent.Render()`.
3. Bubble Tea’s `Update()` handles key/mouse events you wire via component options.
4. Use `SetSize()` during `WindowSizeMsg` to adapt layouts.

---

## ⚙️ CLI Scaffold (`cmd/scaffold`)

* `strawberryui scaffold component <name>`

  * Generates `pkg/components/<name>/<name>.go` with template stubs.
* `strawberryui scaffold theme <name>`

  * Generates a new theme file under `pkg/themes`.

---

## ✅ Testing & CI

* **Unit tests** for each component under `pkg/components/...`

  * Render snapshot tests compare `Render()` output strings.
* **Lint & Format** via `golangci-lint` and `gofmt`.
* **CI Pipeline** (`.github/workflows/ci.yml`):

  * `go test ./...`
  * `golangci-lint run`

---

## 📚 Documentation

* **Getting Started**

  * Installation: `go get github.com/you/strawberryui`
  * Basic “Hello, World!” Bubble Tea + StrawberryUI button sample.
* **API Reference**

  * Auto-generated via `godoc`, linked from README.
* **Theming Guide**

  * Walkthrough on customizing colors, spacing, and creating dark/light variants.
* **Examples**

  * Complete apps in `/examples` demonstrating integration patterns.

---

## 🚀 Roadmap & Next Steps

1. **MVP**: Button, TextInput, List, Table, Dialog + themes + docs.
2. **v0.2**: Export more Bubbles components (progress bars, spinners) wrapped in StrawberryUI styling.
3. **v1.0**: Stable API, theme marketplace (community-contributed themes), GitHub Pages docs.
