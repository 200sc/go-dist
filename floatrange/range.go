package floatrange

// Range represents a range of floating point numbers
type Range interface {
	Poll() float64
	Mult(f float64) Range
	InRange(f float64) bool
	EnforceRange(f float64) float64
}
