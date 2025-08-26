package ui

// NodeOption mutates construction-time properties for a Node.
type NodeOption func(*nodeBase)

// WithChildren sets child nodes.
func WithChildren(children ...Node) NodeOption {
	return func(nb *nodeBase) { nb.kids = append([]Node(nil), children...) }
}

// WithAttr attaches cell-level attributes a renderer can use by convention.
func WithAttr(a Attr) NodeOption {
	return func(nb *nodeBase) { nb.Props()["attr"] = a }
}

// WithPadding sets box padding in cells (T, R, B, L).
func WithPadding(pad struct{ T, R, B, L int }) NodeOption {
	return func(nb *nodeBase) { nb.Props()["padding"] = pad }
}

// WithRadius sets visual corner radius (renderer decides how to realize it).
func WithRadius(r int) NodeOption {
	return func(nb *nodeBase) { nb.Props()["radius"] = r }
}

// WithSize hints preferred size (W,H). Renderer/layout may override.
func WithSize(w, h int) NodeOption {
	return func(nb *nodeBase) { nb.Props()["w"], nb.Props()["h"] = w, h }
}

// WithFlex sets flex grow/shrink/basis (used by your layout engine).
func WithFlex(grow, shrink, basis int) NodeOption {
	return func(nb *nodeBase) {
		nb.Props()["grow"], nb.Props()["shrink"], nb.Props()["basis"] = grow, shrink, basis
	}
}

// WithProp sets an arbitrary prop (escape hatch).
func WithProp(key string, v any) NodeOption {
	return func(nb *nodeBase) { nb.Props()[key] = v }
}
