package pointrange

import "github.com/oakmound/oak/alg/floatgeom"

// NewConstant returns a range which will always return the input constant
func NewConstant(i floatgeom.Point3) Range {
	return Constant(i)
}

// Constant implements Range as a poll
// which always returns the same floatgeom.Point3eger.
type Constant floatgeom.Point3

// Poll returns c cast to an floatgeom.Point3
func (c Constant) Poll() floatgeom.Point3 {
	return floatgeom.Point3(c)
}

// EnforceRange on a constant must return the constant
func (c Constant) EnforceRange(floatgeom.Point3) floatgeom.Point3 {
	return floatgeom.Point3(c)
}

// InRange is only valid for a constant if c == i
func (c Constant) InRange(i floatgeom.Point3) bool {
	return floatgeom.Point3(c) == i
}
