package theme

// Theme bundles Tokens and the derived Styles.
type Theme struct {
	Tokens Tokens
	Styles Styles
	Name   string
}

// New creates a theme from a token set.
func New(name string, tokens Tokens) Theme {
	return Theme{Name: name, Tokens: tokens, Styles: BuildStyles(tokens)}
}
