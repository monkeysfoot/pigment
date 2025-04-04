package cmyk

func NewCMYK(c, m, y, k float64) *Color {
	return &Color{c, m, y, k}
}

func NewHexCMYK(hexstr string) (*Color, error) {
	ret := Color{}
	if err := ret.SetHex(hexstr); err != nil {
		return nil, err
	}
	return &ret, nil
}

func MustHexCMYK(hexstr string) *Color {
	c, err := NewHexCMYK(hexstr)
	if err != nil {
		panic(err)
	}
	return c
}
