package theme

import (
	"strconv"
	"strings"
)

// StyleSpec is a neutral, serializable style delta produced by class strings.
type StyleSpec struct {
	// Colors can be token names ("pink-50") or hex ("#FFCAD4")
	FGHex, BGHex, BorderHex       string
	FGToken, BGToken, BorderToken string

	Bold, Underline *bool

	// Padding (cells)
	P, Px, Py, Pt, Pr, Pb, Pl *int

	// Corners
	RadiusKey string // "none"|"sm"|"md"|"lg"
	Radius    *int   // explicit override
}

// ParseClass converts Tailwind-like utilities into a StyleSpec.
func ParseClass(s string) StyleSpec {
	var spec StyleSpec
	for _, tok := range strings.Fields(s) {
		switch {
		case strings.HasPrefix(tok, "bg-"):
			v := tok[3:]; setColor(&spec.BGToken, &spec.BGHex, v)
		case strings.HasPrefix(tok, "fg-"):
			v := tok[3:]; setColor(&spec.FGToken, &spec.FGHex, v)
		case strings.HasPrefix(tok, "border-"):
			v := tok[7:]; setColor(&spec.BorderToken, &spec.BorderHex, v)

		case tok == "bold":
			b := true; spec.Bold = &b
		case tok == "underline":
			u := true; spec.Underline = &u

		// padding
		case strings.HasPrefix(tok, "px-"):
			n := atoi(tok[3:]); spec.Px = &n
		case strings.HasPrefix(tok, "py-"):
			n := atoi(tok[3:]); spec.Py = &n
		case strings.HasPrefix(tok, "pt-"):
			n := atoi(tok[3:]); spec.Pt = &n
		case strings.HasPrefix(tok, "pr-"):
			n := atoi(tok[3:]); spec.Pr = &n
		case strings.HasPrefix(tok, "pb-"):
			n := atoi(tok[3:]); spec.Pb = &n
		case strings.HasPrefix(tok, "pl-"):
			n := atoi(tok[3:]); spec.Pl = &n
		case strings.HasPrefix(tok, "p-"):
			n := atoi(tok[2:]); spec.P = &n

		// radius
		case strings.HasPrefix(tok, "rounded"):
			parts := strings.SplitN(tok, "-", 2)
			if len(parts) == 2 { spec.RadiusKey = parts[1] } else { spec.RadiusKey = "md" }
		}
	}
	return spec
}

func setColor(tok *string, hex *string, v string) {
	if strings.HasPrefix(v, "#") { *hex = v } else { *tok = v }
}
func atoi(s string) int { n, _ := strconv.Atoi(s); return n }
