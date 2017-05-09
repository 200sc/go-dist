package intrange

// Range represents the ability
// to poll a struct and return an integer,
// distributed over some range dependant
// on the implementing struct.
type Range interface {
	Poll() int
	Mult(int) Range
}
