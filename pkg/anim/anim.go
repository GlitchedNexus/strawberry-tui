// Package anim provides reusable building blocks for time-based animations
// in Bubble Tea UIs: an Animator with easing and lerp helpers.
package anim

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Easing maps a linear progress t in [0,1] into an eased value.
type Easing func(t float64) float64

// Easing functions.
func EaseLinear(t float64) float64    { if t<0 {t=0}; if t>1 {t=1}; return t }
func EaseOutCubic(t float64) float64  { if t<0 {t=0}; if t>1 {t=1}; u:=1-t; return 1 - u*u*u }
func EaseInOutQuad(t float64) float64 { if t<0 {t=0}; if t>1 {t=1}; if t<0.5 {return 2*t*t}; return -1 + (4-2*t)*t }
func EaseOutQuad(t float64) float64   { if t<0 {t=0}; if t>1 {t=1}; return 1 - (1-t)*(1-t) }
func EaseOutBack(t float64) float64   { if t<0 {t=0}; if t>1 {t=1}; c1:=1.70158; c3:=c1+1; return 1 + c3*((t-1)*(t-1)*(t-1)) + c1*((t-1)*(t-1)) }

// Config controls Animator behavior.
type Config struct {
	Duration time.Duration // total duration 0→1
	FPS      int           // frames per second (24–30 is fine)
	Easing   Easing        // easing function (defaults to EaseOutCubic)
}

// Animator tracks an animation progress over time.
type Animator struct {
	cfg      Config
	progress float64
	started  bool
	lastStep time.Time
}

// New creates an Animator with defaults.
func New(cfg Config) *Animator {
	if cfg.Duration <= 0 { cfg.Duration = 200 * time.Millisecond }
	if cfg.FPS <= 0 { cfg.FPS = 30 }
	if cfg.Easing == nil { cfg.Easing = EaseOutCubic }
	return &Animator{cfg: cfg}
}

// Restart resets progress to 0 and starts ticking.
func (a *Animator) Restart() { a.progress = 0; a.started = true; a.lastStep = time.Now() }

// JumpToEnd snaps to 1 and stops.
func (a *Animator) JumpToEnd() { a.progress = 1; a.started = false }

// Running indicates whether the animation is in progress.
func (a *Animator) Running() bool { return a.started && a.progress < 1 }

// Value returns eased progress in [0,1].
func (a *Animator) Value() float64 { return a.cfg.Easing(a.progress) }

// LinearValue returns raw progress in [0,1].
func (a *Animator) LinearValue() float64 { return a.progress }

// Advance updates progress based on elapsed time; call on each Tick.
func (a *Animator) Advance() {
	if !a.started || a.progress >= 1 { return }
	now := time.Now()
	var dt time.Duration
	if a.lastStep.IsZero() { dt = time.Second / time.Duration(a.cfg.FPS) } else { dt = now.Sub(a.lastStep) }
	a.lastStep = now
	step := float64(dt) / float64(a.cfg.Duration)
	a.progress += step
	if a.progress >= 1 { a.progress = 1; a.started = false }
}

// Tick returns a Bubble Tea command to schedule the next frame.
func (a *Animator) Tick() tea.Cmd {
	interval := time.Second / time.Duration(a.cfg.FPS)
	return tea.Tick(interval, func(t time.Time) tea.Msg { return t })
}

// Lerp helpers

// LerpFloat linearly interpolates a→b by t (clamped).
func LerpFloat(a, b, t float64) float64 {
	if t<0 {t=0}; if t>1 {t=1}
	return a + (b-a)*t
}

// LerpInt rounds LerpFloat.
func LerpInt(a, b int, t float64) int { return int(LerpFloat(float64(a), float64(b), t)+0.5) }

// LerpHexRGB interpolates two #RRGGBB hex colors by t and returns #RRGGBB.
func LerpHexRGB(aHex, bHex string, t float64) string {
	r1,g1,b1 := hexToRGB(aHex)
	r2,g2,b2 := hexToRGB(bHex)
	r := int(LerpFloat(float64(r1), float64(r2), t))
	g := int(LerpFloat(float64(g1), float64(g2), t))
	b := int(LerpFloat(float64(b1), float64(b2), t))
	return fmt.Sprintf("#%02x%02x%02x", clamp(r), clamp(g), clamp(b))
}

func hexToRGB(hex string) (r, g, b int) {
	if len(hex) > 0 && hex[0] == '#' { hex = hex[1:] }
	switch len(hex) {
	case 3:
		fmt.Sscanf(hex, "%1x%1x%1x", &r, &g, &b); return r*17, g*17, b*17
	case 6:
		fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b); return r, g, b
	default:
		return 0,0,0
	}
}

func clamp(v int) int { if v<0 {return 0}; if v>255 {return 255}; return v }
