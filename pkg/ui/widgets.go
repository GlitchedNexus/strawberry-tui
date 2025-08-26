package ui

// Text is a leaf node; renderer reads "text" + "attr" from Props.
type textNode struct{ nodeBase }

func Text(id, content string, a Attr, opts ...NodeOption) Node {
	nb := nodeBase{id: NodeID(id), prop: map[string]any{"text": content, "attr": a}}
	for _, opt := range opts { opt(&nb) }
	return &textNode{nb}
}

// Box is a container; padding/attr/radius/size/flex are read from Props.
// Children define its content.
type boxNode struct{ nodeBase }

func Box(id string, opts ...NodeOption) Node {
	nb := nodeBase{id: NodeID(id)}
	for _, opt := range opts { opt(&nb) }
	return &boxNode{nb}
}
