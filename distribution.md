# 📓 Documentation & Versioning Specification for StrawberryUI

This document outlines the technical approach, structure, and key features for documentation, command‑reference, interactive demos, update management, and versioning of StrawberryUI.

---

## 🎯 1. Interactive TUI Demo App (Component Explorer)

### Purpose

Provide an in‑terminal storefront allowing users to browse, preview, and interact with all StrawberryUI components and themes, similar to \[terminal.shop].

### Key Features

* **Category Navigation**: Sidebar with component groups (Buttons, Inputs, Tables, Modals, etc.)
* **Live Previews**: On selection, render a live, interactive instance of the component
* **Theme Toggle**: Switch between Light/Dark/Custom themes on the fly
* **Version Switcher**: Select any released version to see how components render
* **Search**: Incremental search for component names and props
* **Help Panel**: Shows relevant CLI commands (e.g. `strawberryui add button`)

### Implementation

1. **Bubble Tea Model**: fetches manifest (JSON) of components & metadata at startup
2. **Component Renderer**: loads the chosen component package dynamically (via Go plugins or embed)
3. **Theming & Sizing**: use `core.SetTheme` and `SetSize` on resize messages
4. **Version Manifest**: read `versions.json` to populate the version switcher
5. **CLI Skeleton**: expose `strawberryui demo` to launch the app

---

## 🔄 2. CLI Update Checker & Component Updater

### Requirements

* Check for updates to StrawberryUI core and individual components on each CLI invocation
* Offer to install / update to the latest compatible version
* Work offline if no network is available (skip with warning)

### Workflow

1. On any `strawberryui ...` command:

   * Read local `~/.strawberryui/registry.json` with installed versions
   * Query remote `https://api.strawberryui.dev/registry` for `latest` and `tags`
2. **Version Comparison**: use SemVer rules to detect newer `MAJOR`, `MINOR`, `PATCH`
3. **Prompt**: if update found, display changelog excerpt and ask:

   > “A new version 1.2.0 is available. Update now? \[Y/n]”
4. On consent, download the updated component code or CLI binary and place in `~/.strawberryui`
5. Expose commands:

   * `strawberryui update`: manual update
   * `strawberryui versions`: list installed vs remote

---

## 📜 3. Command Reference & Bootstrap Guidance

### CLI Commands

```sh
# Initialize a new project
strawberryui init [--template=basic|with-auth]

# Add components
types: button, input, table, dialog
strawberryui add <component> [--theme=dark]

# List available components & versions
strawberryui list [--remote]

# Change theme for existing project
strawberryui theme set <light|dark|<custom-json-path>>
```

### In‑App Help

* The demo TUI’s footer & help panel automatically display the exact command you need for the current action
* Docs pages will include copy‑to‑clipboard command snippets

---

## 🏷️ 4. Versioning Strategy

* **Semantic Versioning (SemVer)**: `MAJOR.MINOR.PATCH`

  * **MAJOR**: breaking API changes
  * **MINOR**: new features, backwards‑compatible
  * **PATCH**: bug fixes
* **Git Tags & Releases**: each release tagged in GitHub; attach a `CHANGELOG.md` entry
* **Changelogs**:

  * Maintain a `CHANGELOG.md` at repo root
  * Use [Keep a Changelog](https://keepachangelog.com/) format
* **Monorepo & Component Workspaces** (optional):

  * Use Go modules with subdirectories (e.g. `components/button/v1.0.0`) if per‑component versioning is needed
* **Deprecation Policy**:

  * Announce breaking changes one major version ahead in docs
  * Provide codemods or upgrade guides in `docs/migration.md`

---

## 🌐 5. Web‑Based Demo & Documentation Site

### Purpose

Offer a browser‑hosted version of the TUI demo and full docs for users who prefer a GUI.

### Tech Stack

* **Next.js** (React + MDX) or **Gatsby**
* Serve static component previews via \[Xterm.js] or embed dummies
* Host on **Vercel** / **GitHub Pages**

### Features

* **Component Gallery**: visually render each component in React, with live prop knobs
* **Version Toggle**: dropdown to switch documentation and gallery to prior releases (pulls MDX from Git tags)
* **CLI Command Snippets**: “Copy” buttons for each command example
* **Search**: full-text search across API reference & changelog

### CI/CD

* On merge to `main`, auto-deploy docs site
* On GitHub release, generate versioned doc routes (`/v1.2.0/button`)

---

## ⚙️ 6. Reliability & Quality Enhancements

* **Documentation Linting**: MDX lint and spellcheck on CI
* **Automated Snapshot Tests**: compare both TUI demo output & web gallery renders
* **Telemetry (opt‑in)**: track errors in CLI (anonymous) to improve UX
* **Localization**: design docs & CLI messages for future i18n support

---

*With this separate spec, StrawberryUI will offer first‑class documentation, interactive demos, robust version management, and seamless onboarding for both terminal and browser users.*
