package hsv

func NewHSV(h, s, v float64) *Color {
	return &Color{h, s, v}
}

func NewHexHSV(hexstr string) (*Color, error) {
	ret := Color{}
	if err := ret.SetHex(hexstr); err != nil {
		return nil, err
	}
	return &ret, nil
}

func MustHexHSV(hexstr string) *Color {
	c, err := NewHexHSV(hexstr)
	if err != nil {
		panic(err)
	}
	return c
}
