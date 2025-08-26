package theme

// Tokens are raw design values. No Lipgloss here.
type Tokens struct {
	Colors struct {
		Bg, Surface, Text string
		Primary, PrimaryFg string
	}
	Space  map[string]int        // spacing scale (cells)
	Radius map[string]int        // rounded corners (cells)
	Border struct{ Normal, Focused string }
	Motion Motion                // defined in motion.go
}

func DefaultTokens() Tokens {
	var t Tokens

	// Your defaults
	t.Colors.Bg        = "#FFFFFF" // 1
	t.Colors.Surface   = "#FFCAD4" // 2
	t.Colors.Text      = "#242423" // 5
	t.Colors.Primary   = "#F4ACB7" // 3
	t.Colors.PrimaryFg = "#3f0d12" // 4

	t.Space  = map[string]int{"xs":1, "sm":2, "md":4, "lg":6, "xl":8}
	t.Radius = map[string]int{"none":0, "sm":0, "md":1, "lg":2}
	t.Border.Normal  = t.Colors.Text
	t.Border.Focused = t.Colors.PrimaryFg

	// Motion defaults come from motion.go
	t.Motion = DefaultMotion()
	return t
}

// Optional helpers (token lookups)
func (t Tokens) Color(name string) string {
	switch name {
	case "white": return "#FFFFFF"
	case "pink-50": return "#FFCAD4"
	case "pink-60": return "#F4ACB7"
	case "maroon-90": return "#3f0d12"
	case "graphite-90": return "#242423"
	}
	// allow direct pass-through of unknowns
	return name
}
func (t Tokens) SpaceVal(key string) int  { if v,ok:=t.Space[key]; ok {return v}; return 0 }
func (t Tokens) RadiusVal(key string) int { if v,ok:=t.Radius[key]; ok {return v}; return 0 }
