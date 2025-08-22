package theme

import "time"

// Motion defines timing tokens for animations/transitions.
type Motion struct {
	Fast   time.Duration
	Normal time.Duration
	Slow   time.Duration
}

// DefaultMotion returns snappy defaults for terminals.
func DefaultMotion() Motion {
	return Motion{
		Fast:   140 * time.Millisecond,
		Normal: 180 * time.Millisecond,
		Slow:   240 * time.Millisecond,
	}
}
