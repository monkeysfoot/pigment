package rgb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColorRGB_Basic(t *testing.T) {

	t.Run("Test color default", func(t *testing.T) {
		col := &Color{}
		r, g, b := col.RGB()
		assert.Equal(t, uint8(0), r)
		assert.Equal(t, uint8(0), g)
		assert.Equal(t, uint8(0), b)
	})

	coltestdat := []struct {
		r, g, b       uint8
		exr, exg, exb uint8
		exhex         string
	}{
		{0, 0, 0, 0, 0, 0, "#000000"},
		{255, 255, 255, 255, 255, 255, "#ffffff"},
		{255, 0, 0, 255, 0, 0, "#ff0000"},
		{0, 255, 0, 0, 255, 0, "#00ff00"},
		{0, 0, 255, 0, 0, 255, "#0000ff"},
		{0, 128, 128, 0, 128, 128, "#008080"},
	}

	for _, tt := range coltestdat {

		col := &Color{}
		col.SetRGB(tt.r, tt.g, tt.b)
		r, g, b := col.RGB()
		assert.Equal(t, tt.exr, r)
		assert.Equal(t, tt.exg, g)
		assert.Equal(t, tt.exb, b)
		assert.Equal(t, tt.exhex, col.Hex())

		hextst := &Color{}
		_ = hextst.SetHex(tt.exhex)
		assert.Equal(t, tt.exr, hextst.R)
		assert.Equal(t, tt.exg, hextst.G)
		assert.Equal(t, tt.exb, hextst.B)
	}
}
