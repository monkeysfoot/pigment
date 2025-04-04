package rgb

import (
	"fmt"
	"strings"
)

// Color represents a color in the RGB color space.
type Color struct {
	R uint8
	G uint8
	B uint8
}

func (c *Color) RGB() (r, g, b uint8) {
	return c.R, c.G, c.B
}

func (c *Color) SetRGB(r, g, b uint8) {
	c.R = r
	c.G = g
	c.B = b
}

func (c *Color) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

func (c *Color) SetHex(hex string) error {
	hex = strings.TrimPrefix(strings.ToLower(hex), "#")
	if len(hex) != 6 {
		return fmt.Errorf("invalid hex string")
	}
	_, err := fmt.Sscanf(hex, "%02x%02x%02x", &c.R, &c.G, &c.B)
	return err
}
