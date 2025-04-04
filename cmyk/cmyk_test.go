package cmyk

import (
	"github.com/monkeysfoot/pigment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ColorsTest(t *testing.T, col pigment.Colorer, expr, expg, expb uint8) {
	r, g, b := col.RGB()

	//conversion between color spaces are very prone to rounding issues
	const delta = 1.0
	assert.InDelta(t, r, expr, delta)
	assert.InDelta(t, g, expg, delta)
	assert.InDelta(t, b, expb, delta)
}

func ColorInterfaceGetTest[K pigment.Colorer](t *testing.T, col pigment.Colorer, expr, expg, expb uint8, expHex string, newK func() K) {
	ColorsTest(t, col, expr, expg, expb)

	tstrgb := newK()
	tstrgb.SetRGB(expr, expg, expb)
	ColorsTest(t, tstrgb, expr, expg, expb)
	assert.Equal(t, expHex, tstrgb.Hex())

	tsthex := newK()
	err := tsthex.SetHex(expHex)
	assert.NoError(t, err)
	ColorsTest(t, tsthex, expr, expg, expb)
}

func TestColorCMYK_Basic(t *testing.T) {

	t.Run("Test color default", func(t *testing.T) {
		col := &Color{}
		r, g, b := col.RGB()
		//important: Default CMYK color state is white, not black...
		assert.Equal(t, uint8(255), r)
		assert.Equal(t, uint8(255), g)
		assert.Equal(t, uint8(255), b)
	})

	coltestdat := []struct {
		c, m, y, k    float64
		exr, exg, exb uint8
		exhex         string
	}{
		{0, 0, 0, 1, 0, 0, 0, "#000000"},
		{0, 0, 0, 0, 255, 255, 255, "#ffffff"},
		{0, 1, 1, 0, 255, 0, 0, "#ff0000"},
		{1, 0, 1, 0, 0, 255, 0, "#00ff00"},
		{1, 1, 0, 0, 0, 0, 255, "#0000ff"},
		{1, 0, 0, 0.5, 0, 128, 128, "#008080"},
	}

	for _, tt := range coltestdat {
		col := &Color{
			C: tt.c,
			M: tt.m,
			Y: tt.y,
			K: tt.k,
		}

		ColorInterfaceGetTest[*Color](t, col, tt.exr, tt.exg, tt.exb, tt.exhex, func() *Color { return &Color{} })
	}
}
