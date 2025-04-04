package cmyk

import (
	"fmt"
	"github.com/monkeysfoot/pigment"
	"github.com/monkeysfoot/pigment/rgb"
	"math"
)

// Color represents a color in the CMYK color space.
type Color struct {
	C float64
	M float64
	Y float64
	K float64
}

func (c *Color) SetRGB(r, g, b uint8) {
	c.SetFromRGB(rgb.Color{R: r, G: g, B: b})
}

func (c *Color) RGB() (r, g, b uint8) {
	thergb := c.ToRGB()
	return thergb.R, thergb.G, thergb.B
}

func (c *Color) R() uint8 { r, _, _ := c.RGB(); return r }
func (c *Color) G() uint8 { _, g, _ := c.RGB(); return g }
func (c *Color) B() uint8 { _, _, b := c.RGB(); return b }

func (c *Color) Hex() string {
	r, g, b := c.RGB()
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func (c *Color) SetHex(hex string) error {
	var thergb rgb.Color
	err := thergb.SetHex(hex)
	if err != nil {
		return err
	}
	c.SetFromRGB(thergb)
	return nil
}

func (c *Color) SetFromRGB(rgb rgb.Color) {
	r := float64(rgb.R) / 255.0
	g := float64(rgb.G) / 255.0
	b := float64(rgb.B) / 255.0

	k := 1 - math.Max(r, math.Max(g, b))
	c.K = pigment.Clamp01(k)

	if k == 1 {
		c.C, c.M, c.Y = 0, 0, 0
		return
	}

	c.C = (1 - r - k) / (1 - k)
	c.M = (1 - g - k) / (1 - k)
	c.Y = (1 - b - k) / (1 - k)
}

func (c *Color) ToRGB() rgb.Color {
	r := 255 * (1 - pigment.Clamp01(c.C)) * (1 - pigment.Clamp01(c.K))
	g := 255 * (1 - pigment.Clamp01(c.M)) * (1 - pigment.Clamp01(c.K))
	b := 255 * (1 - pigment.Clamp01(c.Y)) * (1 - pigment.Clamp01(c.K))
	return rgb.Color{
		R: uint8(math.Round(r)),
		G: uint8(math.Round(g)),
		B: uint8(math.Round(b)),
	}
}
