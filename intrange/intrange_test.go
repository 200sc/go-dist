package intrange

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// There was a previous implementation of spread
// which was not a different way of initializing linear
// and this test confirmed that they were equivalent.
//
// func TestAreTheyTheSame(t *testing.T) {
// 	rand.Seed(time.Now().UnixNano())

// 	l := NewLinear(0, 6)
// 	bs := NewSpread(3, 3)

// 	lTot := 0
// 	bsTot := 0

// 	for i := 0; i < 100000; i++ {
// 		lTot += l.Poll()
// 		bsTot += bs.Poll()
// 	}

// 	fmt.Println(lTot / 1000)
// 	fmt.Println(bsTot / 1000)

// 	l = l.Mult(2)
// 	bs = bs.Mult(2)

// 	lTot = 0
// 	bsTot = 0

// 	for i := 0; i < 100000; i++ {
// 		lTot += l.Poll()
// 		bsTot += bs.Poll()
// 	}

// 	fmt.Println(lTot / 1000)
// 	fmt.Println(bsTot / 1000)
// }

func TestInfinite(t *testing.T) {
	i := NewInfinite()
	assert.Equal(t, i.Poll(), math.MaxInt32)
	i = i.Mult(5)
	assert.Equal(t, i.Poll(), math.MaxInt32)
}

func TestConstant(t *testing.T) {
	c := NewConstant(5)
	assert.Equal(t, c.Poll(), 5)
	c = c.Mult(5)
	assert.Equal(t, c.Poll(), 25)
}

func TestLinear(t *testing.T) {
	l := NewLinear(0, 10)
	assert.True(t, l.Poll() < 11)
	assert.True(t, l.Poll() > -1)
	l = l.Mult(2)
	assert.True(t, l.Poll() < 21)
	l = NewLinear(10, 0)
	assert.True(t, l.Poll() < 11)
	assert.True(t, l.Poll() > -1)
	l = NewLinear(0, 0)
	assert.Equal(t, l.Poll(), 0)
	l = NewSpread(5, 5)
	assert.True(t, l.Poll() < 11)
	assert.True(t, l.Poll() > -1)
	l = NewSpread(5, -5)
	assert.True(t, l.Poll() < 11)
	assert.True(t, l.Poll() > -1)
	l = NewSpread(5, 0)
	assert.Equal(t, l.Poll(), 5)
}
