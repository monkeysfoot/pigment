package hsl

func NewHSL(h, s, l float64) *Color {
	return &Color{h, s, l}
}

func NewHexHSL(hexstr string) (*Color, error) {
	ret := Color{}
	if err := ret.SetHex(hexstr); err != nil {
		return nil, err
	}
	return &ret, nil
}

func MustHexHSL(hexstr string) *Color {
	c, err := NewHexHSL(hexstr)
	if err != nil {
		panic(err)
	}
	return c
}
