# Theming

All visuals flow from **Tokens → Styles → Theme**.

- `Tokens`: raw design values (hex colors, spacing steps, border colors, motion durations)
- `Styles`: lipgloss `Style`s built from tokens
- `Theme`: bundles tokens + styles and a `Name`

Keep app code free of ad-hoc `lipgloss.NewStyle()`; if a component needs new visuals, add them to `Styles`.
