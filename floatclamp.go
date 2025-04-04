package pigment

import "math"

// Clamp01 restricts a float64 value to the range [0.0, 1.0].
func Clamp01(x float64) float64 {
	if x < 0 {
		return 0
	} else if x > 1 {
		return 1
	}
	return x
}

// ClampDegrees normalizes a float64 representing a float to the range [0.0, 360.0].
func ClampDegrees(x float64) float64 {
	mod := math.Mod(x, 360.0)
	if mod < 0 {
		return mod + 360.0
	}
	return mod
}
