# Theming in StrawberryTUI

## 1. Philosophy

StrawberryTUI separates **design tokens** (raw values) from **styles** (concrete implementations) and **themes** (bundled identities).
This separation makes it easy to:

- enforce consistency across components,
- swap themes at runtime,
- override styles per-instance (utility classes),
- and target both **immediate-mode** (Lipgloss) and **retained-mode** rendering.

---

## 2. Flow of Visuals

All visuals flow through a predictable pipeline:

**Tokens → Styles → Theme → Component**

- **Tokens**: raw design values (colors, spacing, border thickness, radius, motion durations).
- **Styles**: `lipgloss.Style`s and retained-mode `Attr` objects derived from tokens.
- **Theme**: bundles tokens and styles, gives them a `Name`, and provides resolver methods.
- **Component**: consumes preset styles and optionally applies utility overrides (`Class`).

---

## 3. Tokens

```go
type Tokens struct {
  Colors struct {
    Bg, Surface, Text, Primary, PrimaryFg string
  }
  Space   map[string]int  // spacing scale
  Radius  map[string]int  // rounded corners
  Border  struct { Normal, Focused string }
  Motion  Motion
}
```

### Motion Tokens

To keep transitions consistent, motion is also tokenized:

```go
type Motion struct {
  Fast, Normal, Slow time.Duration
}
```

**Defaults**:

- Fast: 140ms
- Normal: 180ms
- Slow: 240ms

Use directly with `pkg/anim`:

```go
a := anim.New(anim.Config{
  Duration: th.Tokens.Motion.Normal,
  FPS:      30,
  Easing:   anim.EaseOutCubic,
})
```

---

## 4. Styles

`Styles` are **derived assets**: ready-to-use style objects pre-built from tokens.

```go
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
```

### Rule:

**Never** call `lipgloss.NewStyle()` directly inside a component.
If a component needs a new visual, extend `Styles` and rebuild them via `BuildStyles(tokens)`.

---

## 5. Theme

A `Theme` bundles both tokens and styles under a single identity:

```go
type Theme struct {
  Name   string
  Tokens Tokens
  Styles Styles
}
```

It also exposes resolvers to bridge between **tokens**, **utility class strings**, and **rendering backends**:

- `ResolveLipgloss(base lipgloss.Style, spec StyleSpec) lipgloss.Style`
- `ResolveTUI(spec StyleSpec) ResolvedTUI`

---

## 6. Utility Classes (Tailwind-like)

Sometimes you need **per-instance overrides** (like Tailwind for the web). StrawberryTUI supports a light grammar:

```
bg-<token|#hex>      → background color
fg-<token|#hex>      → foreground color
border-<token|#hex>  → border color
p-<n>, px-<n>, ...   → spacing (padding)
rounded[-sm|md|lg]   → radius
bold, underline      → text attributes
```

### Example

```go
Button("Save", ButtonOpts{
  Primary: true,
  Class:   "px-4 py-1 rounded bg-#FFCAD4 fg-#3f0d12",
})
```

This string is parsed into a **StyleSpec** (neutral style delta), then resolved by the Theme into either:

- a Lipgloss `Style` (for immediate-mode components), or
- a `ResolvedTUI` (Attr + padding + radius) for the retained-mode renderer.

---

## 7. Precedence of Styles

When combining presets and overrides:

1. **Component defaults** (from `Styles.*`)
2. **Variant overrides** (e.g. `Button.Primary`, `Panel.Header`)
3. **Utility class string** (`Class`)
4. **Stateful overrides** (focused, disabled, active)

This ensures consistency while giving consumers flexibility.

---

## 8. Example Usage

### Immediate-mode (Lipgloss)

```go
style := th.Styles.Button.Base
if opts.Primary {
  style = th.Styles.Button.Primary
}
if opts.Class != "" {
  spec  := theme.ParseClass(opts.Class)
  style = th.ResolveLipgloss(style, spec)
}
return style.Render(opts.Label)
```

### Retained-mode (Node tree)

```go
spec := theme.ParseClass(opts.Class)
res  := th.ResolveTUI(spec)

return ui.Box("btn",
  ui.WithAttr(res.Attr),
  ui.WithPadding(res.Padding),
  ui.WithRadius(res.Radius),
  ui.WithChildren(
    ui.Text("label", opts.Label, res.Attr),
  ),
)
```

---

## 9. Why this Matters

- **Consistency**: tokens define the design system once.
- **Flexibility**: utility classes allow per-instance adjustments without breaking the theme.
- **Dual rendering**: both Lipgloss and the retained-mode renderer consume the same pipeline.
- **Future-proof**: themes can be swapped or extended without rewriting components.

---

## 10. Quick Mental Model

```
Design system → Tokens
Tokens        → Styles
Tokens+Styles → Theme
Theme + Class overrides → Resolved Style
Resolved Style → Lipgloss Style OR Retained-mode Attr
```
