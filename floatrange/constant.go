package floatrange

// Constant is a range that represents some constant float
type Constant float64

// NewConstant returns f wrapped in Constant
func NewConstant(f float64) Range {
	return Constant(f)
}

// Poll returns the float behind the constant
func (c Constant) Poll() float64 {
	return float64(c)
}

// Mult scales c by f
func (c Constant) Mult(f float64) Range {
	c = Constant(float64(c) * f)
	return c
}

// InRange returns whether f == the float behind the constant
func (c Constant) InRange(f float64) bool {
	// Should we use epsilon equality?
	return f == float64(c)
}

// EnforceRange returns the float behind the constant
func (c Constant) EnforceRange(float64) float64 {
	return float64(c)
}

// Percentile returns the float behind the constant
func (c Constant) Percentile(float64) float64 {
	return float64(c)
}
