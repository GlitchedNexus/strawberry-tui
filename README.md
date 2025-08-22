# strawberrytui

A terminal UI component library with **design tokens → styles → components**, plus a small CLI to copy components into your app (like shadcn).

## Quick start

```bash
git clone REPO_URL
cd strawberrytui
go run ./cmd/tui-playground
```

## Try the CLI

```bash
go run ./cmd/tui-cli init ./example-app
go run ./cmd/tui-cli add button ./example-app/ui/button
```

This will materialize themed component code into your app.
