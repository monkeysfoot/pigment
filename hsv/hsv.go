package hsv

import (
	"fmt"
	"github.com/monkeysfoot/pigment"
	"github.com/monkeysfoot/pigment/rgb"
	"math"
)

// Color represents a color in the HSV color space.
type Color struct {
	H float64
	S float64
	V float64
}

func (clr *Color) SetRGB(r, g, b uint8) {
	clr.SetFromRGB(rgb.Color{R: r, G: g, B: b})
}

func (clr *Color) RGB() (r, g, b uint8) {
	c := clr.ToRGB()
	return c.R, c.G, c.B
}

func (clr *Color) R() uint8 { r, _, _ := clr.RGB(); return r }
func (clr *Color) G() uint8 { _, g, _ := clr.RGB(); return g }
func (clr *Color) B() uint8 { _, _, b := clr.RGB(); return b }

func (clr *Color) Hex() string {
	r, g, b := clr.RGB()
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func (clr *Color) SetHex(hex string) error {
	var thecol rgb.Color
	err := thecol.SetHex(hex)
	if err != nil {
		return err
	}
	clr.SetFromRGB(thecol)
	return nil
}

func (clr *Color) SetFromRGB(rgb rgb.Color) {
	r := float64(rgb.R) / 255.0
	g := float64(rgb.G) / 255.0
	b := float64(rgb.B) / 255.0

	maxval := math.Max(r, math.Max(g, b))
	minval := math.Min(r, math.Min(g, b))
	d := maxval - minval

	clr.V = maxval

	if d == 0 {
		clr.H = 0
		clr.S = 0
		return
	}

	clr.S = d / maxval

	switch maxval {
	case r:
		clr.H = (g - b) / d
		if g < b {
			clr.H += 6
		}
	case g:
		clr.H = (b-r)/d + 2
	case b:
		clr.H = (r-g)/d + 4
	}

	clr.H *= 60
	clr.H = math.Mod(clr.H, 360.0)
}

func (clr *Color) ToRGB() rgb.Color {
	h := math.Mod(clr.H, 360.0)
	s := pigment.Clamp01(clr.S)
	v := pigment.Clamp01(clr.V)

	if s == 0 {
		val := uint8(math.Round(v * 255))
		return rgb.Color{R: val, G: val, B: val}
	}

	hi := math.Floor(h / 60)
	f := (h / 60) - hi
	p := v * (1 - s)
	q := v * (1 - f*s)
	t := v * (1 - (1-f)*s)

	var r, g, b float64

	switch int(hi) % 6 {
	case 0:
		r, g, b = v, t, p
	case 1:
		r, g, b = q, v, p
	case 2:
		r, g, b = p, v, t
	case 3:
		r, g, b = p, q, v
	case 4:
		r, g, b = t, p, v
	case 5:
		r, g, b = v, p, q
	}

	return rgb.Color{
		R: uint8(math.Round(r * 255)),
		G: uint8(math.Round(g * 255)),
		B: uint8(math.Round(b * 255)),
	}
}
