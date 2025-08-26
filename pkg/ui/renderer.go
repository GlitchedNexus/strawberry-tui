package ui

// CellOp is a primitive draw operation (renderer-internal but exposed for testing).
type CellOp struct {
	X, Y int
	R    rune
	A    Attr
}

// RenderPlan is what a reconciler/rasterizer produces for Commit.
type RenderPlan struct {
	Ops []CellOp
}

// Engine is the retained-mode renderer contract.
// Your internal engine should satisfy this via an adapter.
type Engine interface {
	Reconcile(prev, next Node, bounds Rect) RenderPlan
	Commit(plan RenderPlan) string // returns frame string for Bubble Tea's View()
}
