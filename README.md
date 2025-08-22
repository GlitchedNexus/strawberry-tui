# strawberrytui

A terminal UI component library with **design tokens → styles → components**, smooth **keyboard-driven animations**, a tiny **anim** utility, and a **CLI** that copies components into your app.

## Features

- Themeable via **design tokens** (colors, spacing, borders, motion)
- Components: **Button**, **Panel**, **SelectList** (animated), **HighlightRow** (variable opacity highlight)
- **pkg/anim**: reusable `Animator` with easing + color/padding lerps
- **Playground** app to preview components
- **CLI** (`tui-cli`) to scaffold theme files and add components (like shadcn)

## Quick start

```bash
git clone https://github.com/GlitchedNexus/strawberrytui
cd strawberrytui
go run ./cmd/tui-playground
```

Use ↑/↓ to move the select list. Press **Tab** to toggle the panel focus glow. Press **h** to toggle highlight on a row.

## CLI (optional)

```bash
# initialize theme skeleton into your app
go run ./cmd/tui-cli init ./myapp

# add a button component into your app
go run ./cmd/tui-cli add button ./myapp/ui/button

# add a select list component into your app
go run ./cmd/tui-cli add selectlist ./myapp/ui/selectlist
```

## Packages

- `pkg/theme` — tokens, styles, motion
- `pkg/anim` — Animator + lerps
- `pkg/components` — button, panel, selectlist, highlightrow

See `docs/` for more guides.
