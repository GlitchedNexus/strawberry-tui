package utils

import (
	"fmt"
)

func hexToRGB(hex string) (r, g, b int) {
	if hex == "" { return 0,0,0 }
	if hex[0] == '#' { hex = hex[1:] }
	switch len(hex) {
	case 3:
		fmt.Sscanf(hex, "%1x%1x%1x", &r, &g, &b)
		r = r*17; g = g*17; b = b*17
	case 6:
		fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
	}
	return
}

func rgbToHex(r,g,b int) string { return fmt.Sprintf("#%02x%02x%02x", clamp(r), clamp(g), clamp(b)) }
func clamp(v int) int { if v<0 {return 0}; if v>255 {return 255}; return v }

// LerpHex linearly interpolates two hex colors by t in [0,1].
func LerpHex(a, b string, t float64) string {
	r1,g1,b1 := hexToRGB(a)
	r2,g2,b2 := hexToRGB(b)
	r := int(float64(r1) + (float64(r2)-float64(r1))*t)
	g := int(float64(g1) + (float64(g2)-float64(g1))*t)
	bb := int(float64(b1) + (float64(b2)-float64(b1))*t)
	return rgbToHex(r,g,bb)
}
