package intrange

// NewConstant returns a range which will always return the input constant
func NewConstant(i int) Range {
	return Constant(i)
}

// Constant implements Range as a poll
// which always returns the same integer.
type Constant int

// Poll returns c cast to an int
func (c Constant) Poll() int {
	return int(c)
}

// Mult returns this range scaled by i
func (c Constant) Mult(i int) Range {
	return Constant(int(c) * i)
}
