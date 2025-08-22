package theme

type Tokens struct {
	Colors struct {
		Bg, Surface, Text, Muted, Primary, PrimaryFg, Warning, Success string
	}
	Space  []int
	Radius []int
	Border struct{ Normal, Focused string }
}

var LightTokens = func() Tokens {
	var t Tokens
	t.Colors = struct {
		Bg, Surface, Text, Muted, Primary, PrimaryFg, Warning, Success string
	}{"#0B0C0F", "#111317", "#E6E8EB", "#9BA3AF", "#3B82F6", "#0B0C0F", "#F59E0B", "#10B981"}
	t.Space = []int{0,1,2,3,4,6,8}
	t.Radius = []int{0,1,2,3}
	t.Border = struct{ Normal, Focused string }{"#2A2F3A", "#3B82F6"}
	return t
}()
