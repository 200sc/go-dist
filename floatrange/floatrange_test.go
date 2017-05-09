package floatrange

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfinite(t *testing.T) {
	i := NewInfinite()
	assert.Equal(t, i.Poll(), math.MaxFloat64)
	i = i.Mult(5)
	assert.Equal(t, i.Poll(), math.MaxFloat64)
	assert.True(t, i.InRange(math.MaxFloat64))
	assert.False(t, i.InRange(3.0))
	assert.Equal(t, i.EnforceRange(3.0), math.MaxFloat64)
}

func TestConstant(t *testing.T) {
	c := NewConstant(5)
	assert.Equal(t, c.Poll(), 5.0)
	c = c.Mult(5)
	assert.Equal(t, c.Poll(), 25.0)
	assert.True(t, c.InRange(25.0))
	assert.False(t, c.InRange(24.0))
	assert.Equal(t, c.EnforceRange(3.0), 25.0)
}

func TestLinear(t *testing.T) {
	l := NewLinear(0, 10)
	assert.True(t, l.Poll() < 11)
	assert.True(t, l.Poll() > -1)
	l = l.Mult(2)
	assert.True(t, l.Poll() < 21)
	l = NewSpread(5, 5)
	assert.True(t, l.Poll() < 11)
	assert.True(t, l.Poll() > -1)
	assert.True(t, l.InRange(3.0))
	assert.False(t, l.InRange(11.0))
	assert.Equal(t, l.EnforceRange(3.0), 3.0)
	assert.Equal(t, l.EnforceRange(-1.0), 0.0)
	assert.Equal(t, l.EnforceRange(11.0), 10.0)
}
