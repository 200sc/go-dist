package floatrange

import "math/rand"

// NewSpread returns a linear range from base-spread to base+spread
func NewSpread(base, spread float64) Range {
	return Linear{base - spread, base + spread, nowRand()}
}

// NewLinear returns a linear range from min to max
func NewLinear(min, max float64) Range {
	return Linear{min, max, nowRand()}
}

// Linear is a range from min to max
type Linear struct {
	Max, Min float64
	rng      *rand.Rand
}

// Poll on a linear float range returns a float at uniform
// distribution in lfr's range
func (lfr Linear) Poll() float64 {
	return ((lfr.Max - lfr.Min) * lfr.rng.Float64()) + lfr.Min
}

// Mult scales a Linear by f
func (lfr Linear) Mult(f float64) Range {
	lfr.Max *= f
	lfr.Min *= f
	return lfr
}

// InRange returns whether an input float could be returned
// by Poll()
func (lfr Linear) InRange(f float64) bool {
	// should still work if min is greater than max
	return (f > lfr.Min) == (f < lfr.Max)
}

// EnforceRange returns f, if InRange(f), or the closest value
// in the range to f.
func (lfr Linear) EnforceRange(f float64) float64 {
	if lfr.InRange(f) {
		return f
	}
	// At this point we need to enforce min < max
	if lfr.Min > lfr.Max {
		lfr.Min, lfr.Max = lfr.Max, lfr.Min
	}
	if f <= lfr.Min {
		return lfr.Min
	}
	return lfr.Max
}

// Percentile returns the fth percentile value along this range
func (lfr Linear) Percentile(f float64) float64 {
	return ((lfr.Max - lfr.Min) * f) + lfr.Min
}
