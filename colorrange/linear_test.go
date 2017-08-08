package colorrange

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinear(t *testing.T) {
	rng := NewLinear(color.RGBA{255, 255, 255, 255}, color.RGBA{255, 255, 255, 255})
	assert.Equal(t, color.RGBA{255, 255, 255, 255}, rng.Poll())
}
