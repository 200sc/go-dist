// Package pointrange holds distributions that accept and return 3d points
package pointrange

import "github.com/oakmound/oak/alg/floatgeom"

// Range represents the ability
// to poll a struct and return an integer,
// distributed over some range dependant
// on the implementing struct.
type Range interface {
	Poll() floatgeom.Point3
	// Mult(float64, float64, float64) Range
	InRange(floatgeom.Point3) bool
	EnforceRange(floatgeom.Point3) floatgeom.Point3
	//Percentile(float64, float64, float64) floatgeom.Point3
}
