package colorrange

import (
	"image/color"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinear(t *testing.T) {
	rng := NewLinear(color.RGBA{255, 255, 255, 255}, color.RGBA{255, 255, 255, 255})
	assert.Equal(t, color.RGBA{255, 255, 255, 255}, rng.Poll())
	for i := 0; i < 100; i++ {
		assert.Equal(t, color.RGBA{255, 255, 255, 255}, rng.Percentile(rand.Float64()))
	}
	rng = NewLinear(color.RGBA{0, 0, 0, 255}, color.RGBA{255, 255, 255, 255})
	for i := 0.0; i < 255; i++ {
		p := i / 255
		uinti := uint8(i)
		assert.Equal(t, color.RGBA{uinti, uinti, uinti, 255}, rng.Percentile(p))
	}
	rng = NewLinear(color.RGBA{255, 255, 255, 255}, color.RGBA{0, 0, 0, 255})
	for i := 255.0; i > 0; i-- {
		p := (255 - i) / 255
		uinti := uint8(i)
		assert.Equal(t, color.RGBA{uinti, uinti, uinti, 255}, rng.Percentile(p))
	}
}
