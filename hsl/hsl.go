package hsl

import (
	"fmt"
	"github.com/monkeysfoot/pigment"
	"math"
	"strings"
)

// Color represents a color in the HSL colorspace.
// Hue (H) is in degrees [0, 360], Saturation (S) and Lightness (L) are in [0.0, 1.0].
type Color struct {
	// H is Hue in degrees [0, 360)
	H float64
	// S Saturation [0, 1]
	S float64
	// L Lightness [0, 1]
	L float64
}

// SetRGB converts an RGB color to HSL and assigns the result to the Color receiver.
func (clr *Color) SetRGB(r, g, b uint8) {
	rf := float64(r) / 255.0
	gf := float64(g) / 255.0
	bf := float64(b) / 255.0

	maxval := math.Max(math.Max(rf, gf), bf)
	minval := math.Min(math.Min(rf, gf), bf)
	l := (maxval + minval) / 2.0

	var h, s float64

	if maxval == minval {
		h = 0
		s = 0
	} else {
		d := maxval - minval
		if l > 0.5 {
			s = d / (2.0 - maxval - minval)
		} else {
			s = d / (maxval + minval)
		}

		switch maxval {
		case rf:
			h = (gf - bf) / d
			if gf < bf {
				h += 6
			}
		case gf:
			h = (bf-rf)/d + 2
		case bf:
			h = (rf-gf)/d + 4
		}
		h /= 6
	}

	clr.H = pigment.ClampDegrees(h * 360.0)
	clr.S = pigment.Clamp01(s)
	clr.L = pigment.Clamp01(l)
}

// Hex returns the hex color string representation (e.g., "#ffcc00") of the HSL color.
func (clr *Color) Hex() string {
	r, g, b := clr.RGB()
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

// SetHex parses a hex color string and sets the HSL values accordingly.
func (clr *Color) SetHex(hex string) error {
	hex = strings.TrimPrefix(strings.ToLower(hex), "#")
	if len(hex) != 6 {
		return fmt.Errorf("invalid hex string")
	}
	var r, g, b uint8
	_, err := fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
	if err != nil {
		return err
	}
	clr.SetRGB(r, g, b)
	return nil
}

// RGB returns the red, green, and blue components of the HSL color.
func (clr *Color) RGB() (r, g, b uint8) {
	h := pigment.ClampDegrees(clr.H) / 360.0
	s := pigment.Clamp01(clr.S)
	l := pigment.Clamp01(clr.L)

	var rF, gF, bF float64

	if s == 0 {
		rF = l
		gF = l
		bF = l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q

		rF = hueToRGB(p, q, h+1.0/3.0)
		gF = hueToRGB(p, q, h)
		bF = hueToRGB(p, q, h-1.0/3.0)
	}
	return uint8(math.Round(rF * 255)), uint8(math.Round(gF * 255)), uint8(math.Round(bF * 255))
}

// hueToRGB is a helper function used to convert HSL to RGB.
func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	switch {
	case t < 1.0/6.0:
		return p + (q-p)*6*t
	case t < 1.0/2.0:
		return q
	case t < 2.0/3.0:
		return p + (q-p)*(2.0/3.0-t)*6
	default:
		return p
	}
}

// R returns the red component of the HSL color.
func (clr *Color) R() uint8 {
	r, _, _ := clr.RGB()
	return r
}

// G returns the green component of the HSL color.
func (clr *Color) G() uint8 {
	_, g, _ := clr.RGB()
	return g
}

// B returns the blue component of the HSL color.
func (clr *Color) B() uint8 {
	_, _, b := clr.RGB()
	return b
}

// HexString is an alias for Hex().
func (clr *Color) HexString() string {
	return clr.Hex()
}
