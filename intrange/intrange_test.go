package intrange

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
