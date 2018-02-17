package pointrange

import (
	"math/rand"

	"github.com/oakmound/oak/alg/floatgeom"
)

// NewLinear returns a linear range
// between min and max
func NewLinear(min, max floatgeom.Point3) Range {
	if max == min {
		return Constant(min)
	}
	min2 := min.LesserOf(max)
	max2 := max.GreaterOf(min)
	return linear{min2, max2, nowRand()}
}

// NewSpread returns a linear range from base - s to base + s
func NewSpread(base, s floatgeom.Point3) Range {
	if s.Magnitude() == 0 {
		return Constant(base)
	}
	if s.Magnitude() < 0 {
		s = s.MulConst(-1)
	}
	return linear{base.Sub(s), base.Add(s), nowRand()}
}

// Linear polls on a linear scale
// between a minimum and a maximum
// linear is private so that the maximum cannot
// be changed to be less than the minimum.
type linear struct {
	Min, Max floatgeom.Point3
	rng      *rand.Rand
}

func (lir linear) Poll() floatgeom.Point3 {
	return lir.Max.Sub(lir.Min).Mul(
		floatgeom.Point3{
			lir.rng.Float64(),
			lir.rng.Float64(),
			lir.rng.Float64(),
		}).Add(lir.Min)
}

func (lir linear) InRange(i floatgeom.Point3) bool {
	if i.X() < lir.Min.X() || i.X() > lir.Max.X() {
		return false
	}
	if i.Y() < lir.Min.Y() || i.Y() > lir.Max.Y() {
		return false
	}
	if i.Z() < lir.Min.Z() || i.Z() > lir.Max.Z() {
		return false
	}
	return true
}

func (lir linear) EnforceRange(i floatgeom.Point3) floatgeom.Point3 {
	return i.GreaterOf(lir.Min).LesserOf(lir.Max)
}
