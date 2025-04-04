package rgb

func NewRGB(r, g, b uint8) *Color {
	return &Color{r, g, b}
}

func NewHexRGB(hexstr string) (*Color, error) {
	ret := Color{}
	if err := ret.SetHex(hexstr); err != nil {
		return nil, err
	}
	return &ret, nil
}

func MustHexRGB(hexstr string) *Color {
	c, err := NewHexRGB(hexstr)
	if err != nil {
		panic(err)
	}
	return c
}
