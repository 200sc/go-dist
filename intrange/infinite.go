package intrange

import "math"

// NewInfinite returns a range which will always return math.MaxInt32 and
// is unchangable.
func NewInfinite() Range {
	return Infinite{}
}

// Infinite is a immutable range which always polls math.MaxInt32
type Infinite struct{}

// Poll returns math.MaxInt32 on Infinites.
func (inf Infinite) Poll() int {
	return math.MaxInt32
}

// Mult does nothing to Infinites.
func (inf Infinite) Mult(i int) Range {
	return inf
}
