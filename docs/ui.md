# pkg/ui — Public API for StrawberryTUI (v0.1)

> **Audience:** component authors and app developers.
> **Purpose:** a small, stable façade over the internal retained‑mode renderer so you can build and compose UI **without** importing engine internals.

---

## 1) Design goals

- **Stable contracts**: insulate components/apps from internal engine changes.
- **Declarative**: build a **Node tree** with composable options (no ad‑hoc maps or random props).
- **Renderer‑agnostic**: the same Node tree can target different backends (ANSI, tcell, tests).
- **Dual‑path**: interoperates with immediate‑mode (Lipgloss) while enabling retained‑mode.

---

## 2) Folder contents

```
pkg/ui/
  types.go       # Color, Attr, Rect (primitive data types)
  node.go        # Node interface + base implementation
  options.go     # Functional options (WithAttr, WithPadding, ...)
  widgets.go     # Convenience constructors: Text, Box
  renderer.go    # Engine (renderer) interface, CellOp/RenderPlan
  adapter.go     # (optional) Adapters to instantiate concrete engines
  doc.go         # Package docs (godoc entrypoint)
```

> **Note:** `adapter.go` is optional; it exposes helpers like `ui.NewANSIEngine(...)` that wrap `internal/renderer` constructors. Keeping adapters here avoids leaking internal packages to app code.

---

## 3) Core types (from `types.go`)

```go
// Color is a terminal color index (-1 = default; map hex→index in the engine).
type Color int16

// Attr are per-cell attributes the renderer understands.
type Attr struct {
    FG, BG   Color
    Bold     bool
    Underline bool
}

// Rect is a cell-space rectangle.
type Rect struct{ X, Y, W, H int }
```

### When to use

- `Attr` is the minimal style payload stored on Nodes (e.g., resolved from theme/class strings).
- `Rect` is used by layout/renderer APIs; app code usually only sets the root bounds.

---

## 4) Nodes (from `node.go` and `widgets.go`)

```go
// Node is the declarative element in the scene graph.
type Node interface {
    ID() NodeID
    Children() []Node
    Props() map[string]any
}

type NodeID string
```

### Constructors

```go
// Text is a leaf node with content and attributes.
func Text(id, content string, a Attr, opts ...NodeOption) Node

// Box is a container node; accepts children and box-like props.
func Box(id string, opts ...NodeOption) Node
```

### Why `Props() map[string]any`?

- It’s a **narrow waist** between public API and engine internals. Public helpers (see below) set well-known keys. Internals are free to optimize representation later without breaking callers.

---

## 5) Functional options (from `options.go`)

```go
// Composition-friendly setters; engine reads well-known keys from Props.
func WithChildren(children ...Node) NodeOption
func WithAttr(a Attr) NodeOption
func WithPadding(pad struct{ T,R,B,L int }) NodeOption
func WithRadius(r int) NodeOption
func WithSize(w, h int) NodeOption          // preferred size hint
func WithFlex(grow, shrink, basis int) NodeOption
func WithProp(key string, v any) NodeOption // escape hatch
```

### Known prop keys (conventions)

These are set by the options above and consumed by the engine:

- `"attr"`: `Attr`
- `"padding"`: struct `{T,R,B,L int}`
- `"radius"`: `int`
- `"w", "h"`: `int` (preferred size hints)
- `"grow", "shrink", "basis"`: `int` (flex layout hints)

> Keep custom keys namespaced (e.g., `"data-role"`, `"aria-label"`) to avoid collisions.

---

## 6) Engine interface (from `renderer.go`)

```go
// CellOp is a primitive draw op (useful for testing & metrics).
type CellOp struct { X, Y int; R rune; A Attr }

type RenderPlan struct { Ops []CellOp }

// Engine is implemented by the retained-mode renderer(s).
type Engine interface {
    Reconcile(prev, next Node, bounds Rect) RenderPlan
    Commit(plan RenderPlan) string // returns frame string for Bubble Tea's View()
}
```

### How apps use it

- Keep the last rendered tree (e.g., in your Bubble Tea model). On each `View()`, build the new tree, call `Reconcile(prev, next, bounds)`, then `Commit(plan)` and return the resulting frame string.

---

## 7) Example: Building a small tree

```go
// Given resolved Attrs from the theme (not shown here):
labelAttr := ui.Attr{Bold: true}
btnAttr   := ui.Attr{}

root := ui.Box("root",
    ui.WithChildren(
        ui.Box("toolbar",
            ui.WithPadding(struct{T,R,B,L int}{0,1,0,1}),
            ui.WithChildren(
                ui.Text("title", "StrawberryTUI", labelAttr),
            ),
        ),
        ui.Box("row",
            ui.WithFlex(1, 1, 0), // take remaining space
            ui.WithChildren(
                ui.Text("btn1", "[ Save ]", btnAttr),
                ui.Text("btn2", "[ Delete ]", btnAttr),
            ),
        ),
    ),
)
```

---

## 8) Example: Engine usage with Bubble Tea

```go
// m.prevTree ui.Node
// m.engine   ui.Engine
// m.w, m.h   ints from WindowSize

func (m model) View() string {
    bounds := ui.Rect{X:0, Y:0, W:m.w, H:m.h}
    next   := buildRootTree(m)              // returns ui.Node
    plan   := m.engine.Reconcile(m.prevTree, next, bounds)
    frame  := m.engine.Commit(plan)
    m.prevTree = next
    return frame
}
```

---

## 9) Working with theming & classes

- Resolve Tailwind-like class strings **outside** `pkg/ui` (in `pkg/theme`).
- Apply resolved values via options:

```go
spec := theme.ParseClass("px-4 py-1 rounded bg-#FFCAD4 fg-#3f0d12")
res  := th.ResolveTUI(spec)
btn  := ui.Box("btn",
    ui.WithAttr(res.Attr),
    ui.WithPadding(res.Padding),
    ui.WithRadius(res.Radius),
    ui.WithChildren(ui.Text("label", "Save", res.Attr)),
)
```

---

## 10) Adapters (`adapter.go`) — optional but convenient

Expose constructors that wrap internal engines without leaking their types:

```go
// func NewANSIEngine(w, h int) (Engine, error)
// func NewTcellEngine(screen tcell.Screen) (Engine, error)
```

This way, app code never imports `internal/renderer`; it only depends on `pkg/ui`.

---

## 11) Best practices

- Treat `NodeID` as **stable** across frames when possible (use semantic IDs).
- Prefer **functional options** over mutating `Props()` directly.
- Keep `Attr` minimal; richer styling should be resolved in `pkg/theme`.
- For custom props, prefix with `data-` if they’re not renderer concerns.
- Use a **single root** node; let layout decide sizes instead of baking coordinates.

---

## 12) FAQ

**Q: Why a `map[string]any` for props?**
A: It decouples public API from internal representation and lets us add/remove keys without version breaks. Public helpers ensure common keys are consistent.

**Q: How do I implement a new widget?**
A: Create a constructor in your component package that returns a `ui.Node` tree. Apply styling via theme resolvers, not ad‑hoc code. Use `WithFlex`/`WithPadding`/`WithAttr` to describe layout and look.

**Q: Can I animate attributes?**
A: Yes. Animation code updates the data that feeds `ui.Node` constructors (e.g., padding, colors). The engine computes diffs and redraws only dirty regions.

---

## 13) Versioning & stability

- `pkg/ui` aims for **semver‑stable** contracts once v1.0 lands. Before then, changes will be documented in `CHANGELOG.md`.
- `internal/renderer` is **not** a public API; it may change at any time.

---

## 14) Testing

- Use a **test engine** implementation that records `RenderPlan.Ops` to assert:

  - number of cells written,
  - regions touched (dirty rects),
  - z‑order invariants.

- Golden‑frame tests: compare committed frame strings for small fixtures.

---

## 15) Glossary

- **Immediate mode**: re-render whole screen each frame (Lipgloss path).
- **Retained mode**: keep a scene graph; diff and update only what changed.
- **Attr**: minimal cell attributes (fg, bg, bold, underline).
- **Node**: declarative element with stable ID, children, and props.
- **Engine**: reconciler + layout + raster + backend commit.
