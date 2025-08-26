package ui

// NodeID is a stable identity used by reconciliation.
type NodeID string

// Node is the declarative element in the scene graph.
// - Immutable across frames (encouraged) for simpler diffing.
// - Props is intentionally generic; keep it small (numbers, strings, bools).
type Node interface {
	ID() NodeID
	Children() []Node
	Props() map[string]any
}

type nodeBase struct {
	id   NodeID
	kids []Node
	prop map[string]any
}

func (n *nodeBase) ID() NodeID          { return n.id }
func (n *nodeBase) Children() []Node    { return n.kids }
func (n *nodeBase) Props() map[string]any {
	if n.prop == nil { n.prop = make(map[string]any) }
	return n.prop
}
