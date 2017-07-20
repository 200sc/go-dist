// Package colorrange holds distributions that accept and return float64s

package colorrange

import (
	"image/color"
)

// Range represents a range of colors
type Range interface {
	Poll() color.Color
	InRange(color.Color) bool
	EnforceRange(color.Color) color.Color
}
