# Stlying & Themes

To centralize theming so that changing, say, the primary color in one place ripples through every component, introduce a single **Theme Registry** that every component pulls its tokens from. Here’s how you can evolve the spec:

---

## 1. Central Theme Registry

### pkg/core/theme.go

```go
package core

// Theme holds all the design tokens for your UI.
type Theme struct {
    PrimaryColor     string
    SecondaryColor   string
    BackgroundColor  string
    ForegroundColor  string
    BorderStyle      lipgloss.Border
    SpacingUnit      int
    // …any other tokens: font weights, shadows, etc.
}

// Default themes:
var Light = Theme{
    PrimaryColor:    "#FF3366",
    SecondaryColor:  "#3366FF",
    BackgroundColor: "#FFFFFF",
    ForegroundColor: "#000000",
    BorderStyle:     lipgloss.RoundedBorder(),
    SpacingUnit:     1,
}

var Dark = Theme{
    PrimaryColor:    "#FF6699",
    SecondaryColor:  "#6699FF",
    BackgroundColor: "#000000",
    ForegroundColor: "#FFFFFF",
    BorderStyle:     lipgloss.DoubleBorder(),
    SpacingUnit:     1,
}

// Current holds the active theme. Components read from this.
var Current = Light

// SetTheme switches the global theme at runtime.
func SetTheme(t Theme) {
    Current = t
}
```

* **Single Source of Truth:** Any time you want to tweak a color, spacing, border style, etc., you only modify `Light` or `Dark` (or your custom `Theme`), and everything reflects that change immediately.
* **Runtime Switching:** Call `core.SetTheme(core.Dark)` in your Bubble Tea `Init()` or on a keypress, then re-render your view to see the new theme applied.

---

## 2. Updating Style Tokens to Read from the Registry

### pkg/core/style.go

```go
package core

import "github.com/charmbracelet/lipgloss"

type StyleToken struct {
    Style lipgloss.Style
}

func NewStyle() StyleToken {
    // start from a blank style
    return StyleToken{Style: lipgloss.NewStyle()}
}

func (s StyleToken) FGPrimary() lipgloss.Style {
    return s.Style.Foreground(Current.PrimaryColor)
}

func (s StyleToken) FGSecondary() lipgloss.Style {
    return s.Style.Foreground(Current.SecondaryColor)
}

func (s StyleToken) BG() lipgloss.Style {
    return s.Style.Background(Current.BackgroundColor)
}

func (s StyleToken) Padding(units int) lipgloss.Style {
    return s.Style.Padding(units * Current.SpacingUnit)
}

func (s StyleToken) Border() lipgloss.Style {
    return s.Style.Border(Current.BorderStyle)
}
```

* **Dynamic Styles:** Whenever you call `.FGPrimary()` you get a style tied to the current theme’s primary color. Swap themes and those calls automatically produce the new color.
* **Composable Tokens:** Combine tokens:

  ```go
  buttonStyle := NewStyle().
      FGPrimary().
      BG().
      Padding(1).
      Border()
  ```

---

## 3. Components Consume the Central Styles

In each component’s `Render()`, replace hard-coded colors/styles with the token helpers:

```go
// pkg/components/button/button.go

func (b *Button) Render() string {
    style := core.NewStyle().
        FGPrimary().
        BG().
        Padding(1).
        Border()

    if b.Focused {
        style = style.Bold(true)
    }

    return style.Render(b.Label)
}
```

Now every button uses the same primary color, border style, and spacing unit defined in `core.Current`.

---

## 4. Theming in Practice

```go
// In your Bubble Tea model…

func (m *model) Init() tea.Cmd {
    // Pick dark mode based on an env var or CLI flag
    if m.useDark {
        core.SetTheme(core.Dark)
    } else {
        core.SetTheme(core.Light)
    }
    return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "d":
            // toggle theme on “d” key
            if core.Current == core.Light {
                core.SetTheme(core.Dark)
            } else {
                core.SetTheme(core.Light)
            }
        }
    }
    return m, nil
}

func (m model) View() string {
    // All your components (buttons, inputs, tables…) will automatically
    // pick up the new colors the next time View() is called.
    return m.rootComponent.Render()
}
```

---

### Summary

1. **One `Theme` struct** holds every token (colors, spacing, borders).
2. **Global var `Current`** is the active theme, switchable at runtime.
3. **StyleToken helpers** read from `Current`, so every component is wired to the same source of truth.

With this in place, updating your primary, secondary, or any other design token in one place instantly propagates throughout StrawberryUI.
