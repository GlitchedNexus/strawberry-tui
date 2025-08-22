package theme

type Theme struct {
	Tokens Tokens
	Styles Styles
	Name   string
}

func New(name string, tokens Tokens) Theme {
	return Theme{Name: name, Tokens: tokens, Styles: BuildStyles(tokens)}
}
