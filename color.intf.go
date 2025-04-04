package pigment

type Colorer interface {
	RGB() (r, g, b uint8)
	SetRGB(r, g, b uint8)
	Hex() string
	SetHex(hex string) error
}
